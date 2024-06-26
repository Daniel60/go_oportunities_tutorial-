package main

import (
	"github.com/Daniel60/go_oportunities_tutorial/config"
	"github.com/Daniel60/go_oportunities_tutorial/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	err := config.Init()
	if err != nil {
		logger.ErrorF("config initialization error: %v", err)
		return
	}
	router.Initilialize()
}
