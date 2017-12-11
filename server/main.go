package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"crypto/worlds/server/state"
)

func main() {
	// TODO: Config/Environment variables
	connectionURL := os.Args[1]
	contractAddress := os.Args[2]

	stateManager := state.NewManager(connectionURL, contractAddress, 5*time.Second)
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
