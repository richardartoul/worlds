package state

import (
	"fmt"
	"log"
	"math"
	"sync"
	"time"

	"crypto/worlds/server/SingleMessage"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Manager is the interface for interacting with the Manager
type Manager interface {
	Init() error
	Get() State
}

// State represents the state of a SingleMessage contract
type State struct {
	Message      string
	PriceInWei   uint64
	PriceInEther float64
}

// NewManager instantiates a new Manager
func NewManager(connAddresses []string, contractAddress string, pollFrequency time.Duration) Manager {
	return &manager{
		RWMutex:         &sync.RWMutex{},
		connAddresses:   connAddresses,
		contractAddress: contractAddress,
		pollFrequency:   pollFrequency,
		conns:           []*ethclient.Client{},
		contracts:       []*SingleMessage.SingleMessage{},
	}
}

type manager struct {
	*sync.RWMutex
	connAddresses   []string
	contractAddress string
	pollFrequency   time.Duration
	conns           []*ethclient.Client
	contracts       []*SingleMessage.SingleMessage
	state           State
}

func (s *manager) Init() error {
	for _, address := range s.connAddresses {
		conn, err := ethclient.Dial(address)
		if err != nil {
			return fmt.Errorf("Failed to connect to the Ethereum client: %v\n", err)
		}
		contract, err := SingleMessage.NewSingleMessage(common.HexToAddress(s.contractAddress), conn)
		if err != nil {
			return fmt.Errorf("Failed instantiating contract: %v\n", err)
		}
		s.conns = append(s.conns, conn)
		s.contracts = append(s.contracts, contract)
	}

	// Fail fast if we can't fetch the contract details right off the bat. All subsequent
	// failed contract details fetches will be log-only.
	err := s.updateState()
	if err != nil {
		return fmt.Errorf("Failed fetching initial contract state: %v\n", err)
	}

	go func() {
		for {
			err := s.updateState()
			if err != nil {
				log.Printf("Error updating state: %v\n", err)
			}
			time.Sleep(s.pollFrequency)
		}
	}()

	return nil
}

func (s *manager) Get() State {
	s.RLock()
	// Copy
	state := s.state
	s.RUnlock()
	return state
}

func (s *manager) updateState() error {
	updatedSuccessfully := false
	for _, contract := range s.contracts {
		message, err := contract.Message(nil)
		if err != nil {
			log.Printf("Failed to retrieve message: %v\n", err)
			continue
		}
		price, err := contract.PriceInWei(nil)
		if err != nil {
			log.Printf("Failed to rerieve price: %v\n", err)
			continue
		}

		s.Lock()
		s.state = State{
			Message:      message,
			PriceInWei:   price.Uint64(),
			PriceInEther: (float64(price.Uint64())) / math.Pow10(18),
		}
		s.Unlock()
		updatedSuccessfully = true
	}

	if !updatedSuccessfully {
		return fmt.Errorf("Failed to update state from any of the available addresses: %v", s.connAddresses)
	}

	return nil
}
