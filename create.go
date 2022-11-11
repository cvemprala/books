package main

import (
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
)

type CreateBookHandler struct {
	repo *Repository
}

func (h CreateBookHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Body == http.NoBody {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var book Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	book.ID = uuid.New().String()
	h.repo.CreateBook(book)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}
