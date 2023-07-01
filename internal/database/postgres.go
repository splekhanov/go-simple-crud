package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

import (
	"github.com/splekhanov/go-simple-crud/internal/model"
)

var DB *gorm.DB
var err error

func DatabaseConnection() {
	host := "localhost"
	port := "5555"
	dbName := "postgres"
	dbUser := "postgres"
	password := "postgres"
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		dbUser,
		dbName,
		password,
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	DB.AutoMigrate(model.Movie{})
	if err != nil {
		log.Fatal("Error connecting to the database...", err)
	}
	fmt.Println("Database connection successful...")
}
