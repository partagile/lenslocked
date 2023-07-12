package main

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// TL;DR - don't export keys nor the type that's being used for context
//
// using custom type as underlying 'any' type this ensures that no one
// else has access to or inadvertently sets a context outside of your control
type ctxKey string

const (
	favoriteColor ctxKey = "favorite-color"
)

func main() {
	// Read config info from .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("SMTP_HOST")
	portStr := os.Getenv("SMTP_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		panic(err)
	}
	username := os.Getenv("SMTP_USERNAME")
	password := os.Getenv("SMTP_PASSWORD")

}
