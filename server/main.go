package main

import (
	"crypto/worlds/server/state"
	"fmt"
	"log"
	"time"
)

func main() {
	// TODO: Config
	stateManager := state.NewManager("http://localhost:8545", "0xbcfc19ae6b67952fa75b10125458f63086e08ce6", 5*time.Second)
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
