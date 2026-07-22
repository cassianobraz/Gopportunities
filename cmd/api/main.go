package main

import (
	"github.com/cassianobraz/Gopportunities/cmd/initializers"
	"github.com/cassianobraz/Gopportunities/internal/config"
	"github.com/cassianobraz/Gopportunities/internal/router"
)

var logger *config.Logger

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	logger = config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	router.Initialize()
}
