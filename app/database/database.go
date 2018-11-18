package database

import (
	"os"
	"strconv"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

// GetDB returns the database connection object.
func GetDB() *gorm.DB {
	return db
}

// Init establishes a connection to the database.
// the configuration are managed in the ".env" file.
func Init() {
	var (
		err error
	)

	if db, err = gorm.Open(os.Getenv("DB_DIALECT"), os.Getenv("DB_CONNECTION")); err != nil {
		panic(err)
	}
	if err = db.DB().Ping(); err != nil {
		panic(err)
	}

	db.LogMode(true)
	maxIdleConns, _ := strconv.Atoi(os.Getenv("DB_IDLE_CONNS"))
	maxOpenConns, _ := strconv.Atoi(os.Getenv("DB_OPEN_CONNS"))
	db.DB().SetMaxIdleConns(maxIdleConns)
	db.DB().SetMaxOpenConns(maxOpenConns)
}
