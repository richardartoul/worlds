package config

import (
	"encoding/json"
	"io/ioutil"
)

// Config represents the configuration struct
type Config struct {
	HTTPPort                    int      `json:"http_port"`
	SSLDomains                  []string `json:"ssl_domains"`
	EthereumClientAddresses     []string `json:"ethereum_client_addresses"`
	EthereumContractAddress     string   `json:"ethereum_contract_addresses"`
	StateRefreshIntervalSeconds int      `json:"state_refresh_interval_seconds"`
}

// Get returns an instance of the configuration struct
func Get() (Config, error) {
	config := Config{}
	configBytes, err := ioutil.ReadFile("./config.json")
	if err != nil {
		return config, err
	}
	err = json.Unmarshal(configBytes, &config)
	return config, err
}
