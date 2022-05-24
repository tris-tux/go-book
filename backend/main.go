package main

import (
	"log"

	"github.com/gin-gonic/gin"
	aa "github.com/tris-tux/go-book/backend/db"
	"github.com/tris-tux/go-book/backend/handler"
	"github.com/tris-tux/go-book/backend/schema"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dbURL := "postgres://postgres:secret@book-postgres:5432/book"
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	db.AutoMigrate(&schema.Book{})

	bookRepo := aa.NewRepo(db)
	bookPostgres := aa.NewPostgres(bookRepo)
	bookHandler := handler.NewBookHandler(bookPostgres)

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/books", bookHandler.GetBooks)
	v1.GET("/books/:id", bookHandler.GetBook)
	v1.POST("/books", bookHandler.CreateBook)
	v1.PUT("/books/:id", bookHandler.UpdateBook)
	v1.DELETE("/books/:id", bookHandler.DeleteBook)

	router.Run()
}
