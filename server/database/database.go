package database

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"janburzinski.de/unilearnapp/database/models"
)

var DB *gorm.DB

func ConnectDB() {
	var err error

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	dbPortStr := os.Getenv("DB_PORT")
	dbPort := 3306

	if dbPortStr != "" {
		if parsedPort, err := strconv.Atoi(dbPortStr); err == nil {
			dbPort = parsedPort
		}
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	config := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}

	DB, err = gorm.Open(mysql.Open(dsn), config)
	if err != nil {
		log.Fatal("Database connection failed", err)
	}

	err = DB.AutoMigrate(
		&models.User{},
		&models.QuestionSheet{},
		&models.SheetStats{},
		&models.Question{},
		&models.Answer{},
	)

	if err != nil {
		log.Fatal("failed to auto migrate database", err)
	}

	log.Println("successfully connected and migrated the database")
}
