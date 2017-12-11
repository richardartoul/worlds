package main

import (
	"crypto/worlds/server/state"
	"fmt"
	"log"
	"time"
)

func main() {
	// TODO: Config
	stateManager := state.NewManager("http://localhost:8545", "0x8bc3ece9ec381ce2ee8d321c6c2393a59a0f5f0a", 5*time.Second)
	err := stateManager.Init()
	if err != nil {
		log.Fatalf("Err initializing state manager: %v", err)
	}

	for {
		state := stateManager.Get()
		fmt.Println("message: ", state.Message)
		fmt.Println("price: ", state.Price)
		time.Sleep(5 * time.Second)
	}
}
