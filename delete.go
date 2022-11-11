package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

type DeleteBookHandler struct {
	repo *Repository
}

func (h DeleteBookHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	ok := h.repo.DeleteBook(id)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
