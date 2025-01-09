package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var Sqldb *sql.DB

func Database_connect() (*sql.DB, error) {
	if err := godotenv.Load(); err != nil {
		log.Printf("Error loading .env file: %v\n", err)
	}
	server := os.Getenv("MSSQL_SERVER")
	user := os.Getenv("MSSQL_USER")
	password := os.Getenv("MSSQL_PASSWORD")
	port := os.Getenv("MSSQL_PORT")
	database := os.Getenv("MSSQL_DATABASE")

	if server == "" || user == "" || port == "" || database == "" {
		log.Printf("MSSQL connection parameters must be set in the environment variables\n")
	}
	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, server, port, database)

	var err error
	Sqldb, err = sql.Open("mysql", connString)
	if err != nil {
		log.Printf("Error connecting to MSSQL: %v\n", err)
		return nil, err
	}

	err = Sqldb.Ping()
	if err != nil {
		log.Printf("Failed to ping MSSQL: %v\n", err)
		return nil, err
	}
	return Sqldb, nil

}
