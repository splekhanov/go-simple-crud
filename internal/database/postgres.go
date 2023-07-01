package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

import (
	"github.com/splekhanov/go-simple-crud/internal/models"
)

var DB *gorm.DB
var err error

type Movie struct {
	gorm.Model
	Title    string `json:"title"`
	Year     string `json:"year"`
	Director string `json:"director"`
	Genre    string `json:"genre"`
	Country  string `json:"country"`
}

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
	DB.AutoMigrate(Movie{})
	if err != nil {
		log.Fatal("Error connecting to the database...", err)
	}
	fmt.Println("Database connection successful...")
}
