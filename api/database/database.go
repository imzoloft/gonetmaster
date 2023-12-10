package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/imzoloft/gonetmaster/api/config"
	"github.com/imzoloft/gonetmaster/logger"
	_ "github.com/lib/pq"
)

var Db *sql.DB

func InitializeDatabase() {
	connectToDatabase()

	createUserTable()
	createAccountTable()
}

func connectToDatabase() {
	db, err := sql.Open("postgres", dsn())

	if err != nil {
		logger.Log.Error("Connecting to database: ", err)
	}

	logger.Log.Info("Connected to database")

	Db = db
}

func createUserTable() {
	ctx, cancelFunc := context.WithTimeout(context.Background(), config.ConnectionTimeout*time.Second)
	defer cancelFunc()

	createUserTableQuery := `
	CREATE TABLE IF NOT EXISTS client (
		id UUID PRIMARY KEY,
		key BYTEA NOT NULL,
		inserted_at TIMESTAMP NOT NULL DEFAULT NOW()
	)`

	_, err := Db.ExecContext(ctx, createUserTableQuery)

	if err != nil {
		logger.Log.Error("Creating user table: ", err)
	}
}

func createAccountTable() {
	ctx, cancelFunc := context.WithTimeout(context.Background(), config.ConnectionTimeout*time.Second)
	defer cancelFunc()

	createAccountTableQuery := `
	CREATE TABLE IF NOT EXISTS account(
		id SERIAL,
		key VARCHAR(64) UNIQUE NOT NULL
	);`
	// Key I used, uuid to base64 to sha256
	// e86b4581-70aa-4416-8dae-6b38c909eddd
	// ZTg2YjQ1ODEtNzBhYS00NDE2LThkYWUtNmIzOGM5MDllZGRk
	// fdcb0a7c658c0835ab597898462e8f64ce6d87c914217e2a5ce7910f3408699d
	_, err := Db.ExecContext(ctx, createAccountTableQuery)

	if err != nil {
		logger.Log.Error("Creating account table: ", err)
	}
}
