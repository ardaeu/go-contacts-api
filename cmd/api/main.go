package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/ardaeu/go-contacts-api/config"
	"github.com/ardaeu/go-contacts-api/internal/handler"
	"github.com/ardaeu/go-contacts-api/internal/storage"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	config.ConnectDB()
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}
	config, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		log.Fatal("pgx config hatası:", err)
	}

	dbpool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatal("pgx pool oluşturulamadı:", err)
	}
	defer dbpool.Close()

	store := storage.NewPGStore()

	h := handler.ContactHandler{Store: store}

	r := chi.NewRouter()

	r.Post("/contacts", h.CreateContact)
	r.Get("/contacts", h.GetAllContacts)
	r.Get("/contacts/{id}", h.GetContactByID)
	r.Put("/contacts/{id}", h.UpdateContact)
	r.Delete("/contacts/{id}", h.DeleteContact)

	log.Println("Sunucu http://localhost:8088 adresinde başladı")
	http.ListenAndServe(":8088", r)
}
