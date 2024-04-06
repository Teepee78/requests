package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/kompere/kompere-api/api"
)

func main() {
	// Load variables from .env file
	if err := godotenv.Load(); err != nil {
		panic(fmt.Sprintf("Error loading .env file: %v", err))
	}

	api.StartServer()
}
