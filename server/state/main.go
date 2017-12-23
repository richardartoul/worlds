package state

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"strconv"
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
	EthPriceUSD  float64
}

type ethereumCoinMarketCapResponse struct {
	PriceUSD string `json:"price_usd"`
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
		httpClient: &http.Client{
			// TODO: Config
			Timeout: 5 * time.Second,
		},
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
	httpClient      *http.Client
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
	err := s.updateContractState()
	if err != nil {
		return fmt.Errorf("Failed fetching initial contract state: %v\n", err)
	}
	err = s.updateEthPriceState()
	if err != nil {
		return fmt.Errorf("Failed fetching ethereum price in USD: %v\n", err)
	}

	// Background goroutine to keep contract state up to date
	go func() {
		for {
			err := s.updateContractState()
			if err != nil {
				log.Printf("Error updating contract state: %v\n", err)
			}
			time.Sleep(s.pollFrequency)
		}
	}()

	// Background goroutine to keep ethereum price up to date
	go func() {
		for {
			err = s.updateEthPriceState()
			if err != nil {
				log.Printf("Error updating ethereum price state: %v\n", err)
			}
			// TODO: Make separate config
			time.Sleep(12 * s.pollFrequency)
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

func (s *manager) updateContractState() error {
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
		// If we made it this far, there is no need to look at the other sources of data
		break
	}

	if !updatedSuccessfully {
		return fmt.Errorf("Failed to update state from any of the available addresses: %v", s.connAddresses)
	}

	return nil
}

func (s *manager) updateEthPriceState() error {
	resp, err := s.httpClient.Get("https://api.coinmarketcap.com/v1/ticker/ethereum/")
	if err != nil {
		return err
	}
	jsonBytes, err := ioutil.ReadAll(io.LimitReader(resp.Body, 1049000))
	if err != nil {
		return err
	}
	ethPriceUSDResp := []*ethereumCoinMarketCapResponse{}
	json.Unmarshal(jsonBytes, ethPriceUSDResp)
	if len(ethPriceUSDResp) != 1 {
		return fmt.Errorf("Received invalid JSON payload: %s", string(jsonBytes))
	}
	ethPriceUSD, err := strconv.ParseFloat(ethPriceUSDResp[0].PriceUSD, 64)
	if err != nil {
		return fmt.Errorf("Received invalid JSON payload: %s, unable to parseFloat: %v", string(jsonBytes), err)
	}

	s.Lock()
	s.state.EthPriceUSD = ethPriceUSD
	s.Unlock()

	return nil
}
