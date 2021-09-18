package config

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

const (
	driver = "sqlite3"
	path   = "./db/stats.sqlite"
	//path   = "/opt/stats.sqlite"
)

// DB Exposed global database connection configurations.
// Manages opening and closing database connections.
// Queries can be executed on this variable.
var DB *sql.DB

// InitDB Loads and builds database connection configurations and writes them into SQL DB variable.
// Causes panic if database metadata are not successfully loaded.
func InitDB() error {
	db, err := sql.Open(driver, path)
	if err != nil {
		panic("Unable to load database meta")
	}
	DB = db
	return err
}
