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
	Discount    int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type BookInput struct {
	Title       string      `json:"title" binding:"required"`
	Price       json.Number `json:"price" binding:"required,number"`
	Description string      `json:"description" binding:"required"`
	Rating      json.Number `json:"rating" binding:"required,number"`
	Discount    json.Number `json:"discount" binding:"required,number"`
}
