package main

import (
	"github.com/joho/godotenv"
	"log"
)

func main() {
	// load env
	err := godotenv.Load()
	if err != nil {
		log.Print("Error loading .env file")
	}
}
