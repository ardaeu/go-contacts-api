package main

import (
	"log"
	"net/http"

	"github.com/ardaeu/contact-api/config"
	"github.com/ardaeu/contact-api/handler"
	"github.com/go-chi/chi/v5"
)

func main() {
	// 1. Veritabanı bağlantısını başlat
	config.ConnectDB()

	// 2. Router oluştur
	r := chi.NewRouter()

	// 3. Contact handler'ı başlat
	contactHandler := handler.NewContactHandler()

	// 4. Route'ları tanımla
	r.Route("/contacts", func(r chi.Router) {
		r.Post("/", contactHandler.CreateContact)
		r.Get("/", contactHandler.GetAllContacts)
		r.Get("/{id}", contactHandler.GetContactByID)
		r.Put("/{id}", contactHandler.UpdateContact)
		r.Delete("/{id}", contactHandler.DeleteContact)
	})

	// 5. Sunucuyu başlat
	log.Println("Sunucu http://localhost:8080 adresinde çalışıyor")
	http.ListenAndServe(":8080", r)
}
