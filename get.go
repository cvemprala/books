package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type GetBooksHandler struct {
	repo *Repository
}

type GetBookHandler struct {
	repo *Repository
}

func (h GetBooksHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h.repo.GetBooks())
}

func (h GetBookHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	book, ok := h.repo.GetBook(id)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}
