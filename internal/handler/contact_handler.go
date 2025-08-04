package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ardaeu/go-contacts-api/internal/model"
	"github.com/ardaeu/go-contacts-api/internal/storage"
	"github.com/go-chi/chi/v5"
)

type ContactHandler struct {
	Store *storage.PGStore
}

func (h *ContactHandler) CreateContact(w http.ResponseWriter, r *http.Request) {
	var c model.Contact
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, "Geçersiz veri", http.StatusBadRequest)
		return
	}

	if err := h.Store.Create(r.Context(), &c); err != nil {
		http.Error(w, "Kayıt oluşturulamadı", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(c)
}

func (h *ContactHandler) GetAllContacts(w http.ResponseWriter, r *http.Request) {
	contacts, err := h.Store.GetAll(r.Context())
	if err != nil {
		http.Error(w, "Veriler alınamadı", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(contacts)
}

func (h *ContactHandler) GetContactByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Geçersiz ID formatı", http.StatusBadRequest)
		return
	}

	contact, err := h.Store.GetByID(r.Context(), id)
	if err != nil {
		if err == storage.ErrNotFound {
			http.NotFound(w, r)
		} else {
			http.Error(w, "Veri alınamadı", http.StatusInternalServerError)
		}
		return
	}

	json.NewEncoder(w).Encode(contact)
}

func (h *ContactHandler) UpdateContact(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Geçersiz ID formatı", http.StatusBadRequest)
		return
	}

	var c model.Contact
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, "Geçersiz veri", http.StatusBadRequest)
		return
	}
	c.ID = id

	if err := h.Store.Update(r.Context(), &c); err != nil {
		if err == storage.ErrNotFound {
			http.NotFound(w, r)
		} else {
			http.Error(w, "Güncellenemedi", http.StatusInternalServerError)
		}
		return
	}

	json.NewEncoder(w).Encode(c)
}

func (h *ContactHandler) DeleteContact(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Geçersiz ID formatı", http.StatusBadRequest)
		return
	}

	if err := h.Store.Delete(r.Context(), id); err != nil {
		if err == storage.ErrNotFound {
			http.NotFound(w, r)
		} else {
			http.Error(w, "Silinemedi", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
