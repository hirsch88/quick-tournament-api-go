package main

import (
	"github.com/hirsch88/quick-tournament-api-go/app"
	"github.com/joho/godotenv"

	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	// Loads the env variables
	godotenv.Load()

	app.Bootstrap()
}
