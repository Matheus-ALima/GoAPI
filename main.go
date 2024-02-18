package main

import (
	"github.com/Matheus-ALima/GoAPI.git/config"
	"github.com/Matheus-ALima/GoAPI.git/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	//initialize config
	err := config.Init()
	if err != nil {
		logger.Errorf("Config initialization error: %v", err)
		return
	}

	//initialize router
	router.Initialize()
}
