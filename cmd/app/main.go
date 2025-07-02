package main

import (
	"fmt"
	"os"
	"rozoomcool/go-api-skeleton/internal/config"
	"rozoomcool/go-api-skeleton/internal/logger"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(fmt.Sprintf("Failed to load configs: %v", err))
	}

	logFile, err := os.OpenFile(cfg.LogFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(fmt.Sprintf("Failed to find log file: %v", err))
	}

	log := logger.New(cfg.Debug, logFile)

	log.Infow("Hello World!")

	fmt.Println(cfg.DBPassword)
}
