package config

import (
	"backend-jual-beli/models"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
)

var JwtSecret []byte

func InitConfig() {
    // Load .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    // Fetch JWT_SECRET from environment variables
    secret := os.Getenv("JWT_SECRET")
    if secret == "" {
        log.Fatalf("JWT_SECRET not found in environment variables")
    }

    // Set JWT secret key
    JwtSecret = []byte(secret)
}

var DB *gorm.DB

// SetupDatabase initializes the database connection and performs migrations
func SetupDatabase() {
	// Database connection string (DSN)
	dsn := "root:@tcp(127.0.0.1:3306)/jual_beli_db?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to MySQL: %v", err)
	}

	log.Println("Connected to the database successfully!")

	// Run migrations to create tables based on models
	if err := DB.AutoMigrate(&models.User{}, &models.Item{}, &models.Order{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	log.Println("Database migration completed!")
}
