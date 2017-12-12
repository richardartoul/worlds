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
func NewManager(connAddress, contractAddress string, pollFrequency time.Duration) Manager {
	return &manager{
		RWMutex:         &sync.RWMutex{},
		connAddress:     connAddress,
		contractAddress: contractAddress,
		pollFrequency:   pollFrequency,
	}
}

type manager struct {
	*sync.RWMutex
	connAddress     string
	contractAddress string
	pollFrequency   time.Duration
	conn            *ethclient.Client
	contract        *SingleMessage.SingleMessage
	state           State
}

func (s *manager) Init() error {
	conn, err := ethclient.Dial(s.connAddress)
	if err != nil {
		return fmt.Errorf("Failed to connect to the Ethereum client: %v\n", err)
	}
	s.conn = conn

	contract, err := SingleMessage.NewSingleMessage(common.HexToAddress(s.contractAddress), conn)
	if err != nil {
		return fmt.Errorf("Failed instantiating contract: %v\n", err)
	}
	s.contract = contract

	// Fail fast if we can't fetch the contract details right off the bat. All subsequent
	// failed contract details fetches will be log-only.
	err = s.updateState()
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
	message, err := s.contract.Message(nil)
	if err != nil {
		return fmt.Errorf("Failed to retrieve message: %v\n", err)
	}
	price, err := s.contract.PriceInWei(nil)
	if err != nil {
		return fmt.Errorf("Failed to rerieve price: %v\n", err)
	}

	s.Lock()
	s.state = State{
		Message:      message,
		PriceInWei:   price.Uint64(),
		PriceInEther: (float64(price.Uint64())) / math.Pow10(18),
	}
	s.Unlock()
	return nil
}
