package main

import (
	"github.com/hirsch88/quick-tournament-api-go/app"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	// Loads the env variables
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal("Error loading .env file")
	}

	app.Bootstrap()
}
