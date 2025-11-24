package app

import (
	"database/sql"
	"fmt"
	"os"
	"time"
	_"github.com/jackc/pgx/v5/stdlib"
	"notes-app/backend/helper"
)

func NewDB() *sql.DB {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)
	db, err := sql.Open("pgx", dsn)
	helper.PanicIfError(err)

	err = db.Ping()
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}
	fmt.Println("Successfully connected to database")

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}