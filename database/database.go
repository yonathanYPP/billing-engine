package database

import (
	"billing-engine/app/models"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB   *gorm.DB
	once sync.Once
)

type Config struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
	SSLMode  string `json:"sslmode"`
}

// LoadConfig membaca file database.json
func LoadConfig() (*Config, error) {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Konversi port ke integer
	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatalf("Invalid port number: %s", os.Getenv("DB_PORT"))
	}

	config := Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     dbPort,
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	return &config, nil
}

// InitDB menginisialisasi koneksi ke PostgreSQL
func InitDB() *gorm.DB {
	config, err := LoadConfig()
	if err != nil {
		log.Fatal("Error loading database config:", err)
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Singleton pattern untuk inisialisasi database hanya sekali
	once.Do(func() {
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("Failed to connect to database:", err)
		}
		log.Println("Database connected successfully!")

		// Auto-migrate semua model
		err = DB.AutoMigrate(GetModels()...)
		if err != nil {
			log.Fatal("Failed to migrate database:", err)
		}
		log.Println("Database migrated successfully!")
	})
	return DB
}

// GetModels mengembalikan semua model yang akan dimigrasi
func GetModels() []interface{} {
	return []interface{}{
		&models.Loan{},
		&models.Payment{},
	}
}
