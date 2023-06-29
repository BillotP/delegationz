package db

import (
	"database/sql"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/jackc/pgx/v4/stdlib"
)

var inst *sql.DB

func Get(dbURL string) *sql.DB {
	if inst != nil {
		return inst
	}
	// Configure the PostgreSQL connection
	connConfig, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		log.Fatal("Error parsing connection config:", err)
	}

	// Create the database connection
	db, err := sql.Open("pgx", connConfig.ConnString())
	if err != nil {
		log.Fatal("Error opening database connection:", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal("Error checking database connection:", err)
	}
	log.Printf("[INFO] Connected to db %s @ %s",
		connConfig.ConnConfig.Database, connConfig.ConnConfig.Host)
	inst = db
	return inst
}
