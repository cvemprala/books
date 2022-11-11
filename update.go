package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type UpdateBookHandler struct {
	repo *Repository
}

func (h UpdateBookHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

	vars := mux.Vars(r)
	id := vars["id"]

	ok := h.repo.UpdateBook(id, book)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}
