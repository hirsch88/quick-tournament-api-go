package main

import (
	"github.com/hirsch88/quick-tournament-api-go/app"
	"github.com/hirsch88/quick-tournament-api-go/app/database"
	"github.com/joho/godotenv"

	_ "github.com/heroku/x/hmetrics/onload"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	// Loads the env variables
	godotenv.Load()

	// Establish a connection to the database and
	// migrate the database schema.
	database.Init()
	database.AutoCreateTables()
	database.AutoMigrateTables()

	// Bootstrap api with all the registered endpoints
	app.Bootstrap()
}
