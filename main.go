package main

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"w3tec.ch/go-heroku/app"
)

func main() {
	// Loads the env variables
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal("Error loading .env file")
	}

	app.Bootstrap()
}
