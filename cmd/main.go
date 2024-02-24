package main

import (
	"github.com/gmessias/api-go-money/database"
	"github.com/gmessias/api-go-money/routes"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.Connect()
	routes.HandleRequests()
}
