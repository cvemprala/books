package main

import (
	"time"
)

type Book struct {
	ID            string    `json:"id"`
	Title         string    `json:"title"`
	Author        string    `json:"author"`
	Price         float32   `json:"price"`
	Publisher     string    `json:"publisher"`
	PublishedDate time.Time `json:"publishedDate"`
}
