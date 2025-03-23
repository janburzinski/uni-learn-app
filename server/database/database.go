package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
)

var DB *sql.DB

/**
* Create and overwrite (DEBUG ONLY) the tables
 */
func SetupDB() (err error) {
	return nil
}

func ConnectDB() {
	var err error

	db_host := os.Getenv("DB_HOST")
	db_user := os.Getenv("DB_USER")
	db_name := os.Getenv("DB_NAME")
	db_password := os.Getenv("DB_PASSWORD")

	db_port_str := os.Getenv("DB_PORT")
	db_port := int64(3306)

	if db_port_str != "" {
		if parsedPort, err := strconv.ParseInt(db_port_str, 10, 64); err == nil {
			db_port = parsedPort
		}
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", db_user, db_password, db_host, db_port, db_name)
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Database connection failed", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Error while pinging database", err)
	}

	if os.Getenv("ENV") == "DEBUG" {
		err = SetupDB()
		if err != nil {
			log.Fatal("error while setting up db", err)
		}
	}
}
