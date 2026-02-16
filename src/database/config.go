package database

import (
	"fmt"
	"log"
	"lppm/src/models"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// Ganti sesuai kredensial Postgres kamu
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	tz := os.Getenv("DB_TIMEZONE")

	// Susun DSN
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		host, user, pass, name, port, tz)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal koneksi ke database:", err)
	}

	// Auto Migration (Membuat tabel otomatis)
	err = database.AutoMigrate(
		&models.User{},
		&models.Admin{},
		&models.Pelanggan{},
		&models.Kategori{},
		&models.Buku{},
		&models.Pesanan{},
		&models.PesananItem{},
	)

	if err != nil {
		log.Fatal("Gagal migrasi database:", err)
	}

	DB = database
	fmt.Println("Database berhasil terkoneksi dan termigrasi!")
}
