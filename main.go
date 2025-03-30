package main

import (
	"aggregator/internal/config"
	"fmt"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Errorf("error reading config", err)
	}

	cfg.SetUser("geert")

	// Reading config again!
	cfg, err = config.Read()
	if err != nil {
		fmt.Errorf("error reading config", err)
	}
	fmt.Printf("Current user name: %v", cfg)
}

