package container

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"payment-platform/config"
	"time"
)

func NewDatabase(config *config.DatabaseConfig) *sql.DB {
	db, err := sql.Open("postgres", config.DSN)
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)

	return db
}
