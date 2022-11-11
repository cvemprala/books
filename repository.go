package main

import (
	"sync"
)

type Repository struct {
	sync.Mutex
	books map[string]Book
}

func NewRepository() *Repository {
	return &Repository{
		books: make(map[string]Book),
	}
}

func (r *Repository) GetBooks() []Book {
	r.Lock()
	defer r.Unlock()

	books := make([]Book, 0, len(r.books))
	for _, book := range r.books {
		books = append(books, book)
	}

	return books
}

func (r *Repository) GetBook(id string) (Book, bool) {
	r.Lock()
	defer r.Unlock()

	book, ok := r.books[id]
	return book, ok
}

func (r *Repository) CreateBook(book Book) {
	r.Lock()
	defer r.Unlock()
	r.books[book.ID] = book
}

func (r *Repository) UpdateBook(id string, book Book) bool {
	r.Lock()
	defer r.Unlock()

	_, ok := r.books[id]
	if !ok {
		return false
	}

	book.ID = id
	r.books[id] = book
	return true
}

func (r *Repository) DeleteBook(id string) bool {
	r.Lock()
	defer r.Unlock()

	_, ok := r.books[id]
	if !ok {
		return false
	}

	delete(r.books, id)
	return true
}
