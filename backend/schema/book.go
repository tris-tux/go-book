package schema

import (
	"encoding/json"
	"time"
)

type Book struct {
	ID          int
	Title       string
	Description string
	Price       int
	Rating      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type BookInput struct {
	Title string `json:"title" binding:"required"`
	// Price int    `json:"price" binding:"required,number"`
	Price json.Number `json:"price" binding:"required,number"`
}
