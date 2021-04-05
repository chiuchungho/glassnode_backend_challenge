package database

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"

	_ "github.com/lib/pq"
)

var db *sql.DB

/*
 * Initialize the database connection
 */
func InitializeSQL() error {
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")
	dbName := os.Getenv("POSTGRES_DB")

	port, err := strconv.ParseUint(dbPort, 10, 32)

	if err != nil {
		log.Error("Error parsing str to int")
		return err
	}

	log.Info("DB Info")
	log.Info(fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbHost, port, dbUser, dbPassword, dbName))

	dBConnection, err := sql.Open("postgres",
		fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbHost, port, dbUser, dbPassword, dbName))

	if err != nil {
		log.Error("Connection Failed!! ",err.Error())
		return err
	}

	db = dBConnection
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(10 * time.Minute)

	log.Info("Connected to db")

	return nil
}

/*
 * Get database connection
 */
func GetSQLConnection() *sql.DB {
	return db
}
