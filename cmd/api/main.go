package main

import (
	"fmt"

	"github.com/cassianobraz/Gopportunities/cmd/initializers"
	"github.com/cassianobraz/Gopportunities/internal/config"
	"github.com/cassianobraz/Gopportunities/internal/router"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	err := config.Init()
	if err != nil {
		fmt.Println(err)
		return
	}

	router.Initialize()
}
