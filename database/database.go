package database

import (
	"aquaculture/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

var (
	DB_USERNAME string = os.Getenv("DB_USERNAME")
	DB_PASSWORD string = os.Getenv("DB_PASSWORD")
	DB_NAME     string = os.Getenv("DB_NAME")
	DB_HOST     string = os.Getenv("DB_HOST")
	DB_PORT     string = os.Getenv("DB_PORT")
)

func InitDB() {
	var err error

	var dsn string = fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s",
		DB_USERNAME,
		DB_PASSWORD,
		DB_HOST,
		DB_PORT,
		DB_NAME,
	)
	pgConfig := postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	})

	DB, err = gorm.Open(pgConfig, &gorm.Config{})

	if err != nil {
		log.Fatalf("error when connecting to the database: %s\n", err)
	}

	log.Println("connected to the database")

	Migrate()
}

func Migrate() {
	err := DB.AutoMigrate(&models.Admin{}, &models.AquacultureFarms{}, &models.Article{}, &models.Farm{}, &models.FarmType{}, &models.Product{}, &models.ProductType{}, &models.Transaction{}, &models.TransactionDetail{}, &models.User{}, &models.PromoCode{}, &models.FarmCondition{})

	if err != nil {
		log.Fatalf("error when migration: %s\n", err)
	}

	log.Println("migration successful")
}
