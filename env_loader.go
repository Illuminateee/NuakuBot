package main

import (
	"log"

	"github.com/joho/godotenv"
)

// EnvLoader handles environment loading
type EnvLoader struct{}

// LoadEnv loads environment variables from a .env file
func (e EnvLoader) LoadEnv() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return nil
}
