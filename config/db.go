package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var DB *pgxpool.Pool

func ConnectDB() {
	var err error

	err = godotenv.Load()
	if err != nil {
		log.Fatal("'.env' dosyası yüklenemedi")
	}

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	DB, err = pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Veritabanı bağlantısı kurulamadı: %v", err)
	}

	err = DB.Ping(context.Background())
	if err != nil {
		log.Fatalf("Veritabanına erişilemiyor: %v", err)
	}

	fmt.Println("Veritabanına bağlanıldı.")
}
