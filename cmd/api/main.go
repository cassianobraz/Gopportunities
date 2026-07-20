package main

import (
	"github.com/cassianobraz/Gopportunities/cmd/initializers"
	"github.com/cassianobraz/Gopportunities/internal/router"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	router.Initialize()
}
