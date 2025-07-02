package main

import (
	"fmt"
	"rozoomcool/go-api-skeleton/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(fmt.Sprintf("Failed to load configs: %v", err))
	}

	fmt.Println(cfg.DBPassword)
}
