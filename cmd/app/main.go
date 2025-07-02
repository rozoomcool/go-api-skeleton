package main

import (
	"fmt"
	"os"
	"rozoomcool/go-api-skeleton/internal/config"
	"rozoomcool/go-api-skeleton/internal/db"
	"rozoomcool/go-api-skeleton/internal/logger"
)

func main() {
	// Init configs
	cfg, err := config.Load()
	if err != nil {
		panic(fmt.Sprintf("Failed to load configs: %v", err))
	}

	// Open logfile
	logFile, err := os.OpenFile(cfg.LogFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(fmt.Sprintf("Failed to find log file: %v", err))
	}

	// Init logger
	log := logger.New(cfg.Debug, logFile)
	defer log.Sync()

	// Init db connection
	dbConn, err := db.Init(cfg)
	if err != nil {
		panic(fmt.Sprintf("Failed initialize database connection: %v", err))
	}
	defer dbConn.Close()

	log.Infow("Hello World!")

	fmt.Println(cfg.DBPassword)
}
