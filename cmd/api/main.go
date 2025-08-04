package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/ardaeu/go-contacts-api/internal/handler"
	"github.com/ardaeu/go-contacts-api/internal/storage"
	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
)

func main() {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgres://postgres:postgres@localhost:5432/go_contacts?sslmode=disable"
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("DB bağlantı hatası: %v", err)
	}
	defer db.Close()

	store := storage.NewPGStore(db)
	h := &handler.ContactHandler{Store: store}

	r := chi.NewRouter()

	r.Post("/contacts", h.CreateContact)
	r.Get("/contacts", h.GetAllContacts)
	r.Get("/contacts/{id}", h.GetContactByID)
	r.Put("/contacts/{id}", h.UpdateContact)
	r.Delete("/contacts/{id}", h.DeleteContact)

	// Sunucuyu başlat
	log.Println("Sunucu http://localhost:8080 adresinde başlatıldı...")
	http.ListenAndServe(":8080", r)
}
