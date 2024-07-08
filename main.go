package main

import (
	"fmt"

	"github.com/anderson89marques/gopportunities/config"
	"github.com/anderson89marques/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	// Initialize Configs
	err := config.Init()

	if err != nil {
		fmt.Println(err)
		logger.Errorf("Config initialization error: %v", err)
		return
	}

	// Initialize Router
	router.Initialize()
}
