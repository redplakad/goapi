package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	DB *sql.DB
}

func NewDatabase(config *Config) *Database {
	dsn := config.GetDSN()
	
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Test the connection
	if err = db.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	fmt.Println("Successfully connected to MySQL database!")

	return &Database{
		DB: db,
	}
}

func (d *Database) Close() {
	if d.DB != nil {
		d.DB.Close()
	}
}
