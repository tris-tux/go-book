package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/tris-tux/go-book/backend/handler"
	"github.com/tris-tux/go-book/backend/schema"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	// DB := db.Init()
	// h := handler.New(DB)

	// dsn := "host=book-postgres user=postgres password=secret dbname=book port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	// _, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// if err != nil {
	// 	log.Fatal("Db connection error")
	// }

	dbURL := "postgres://postgres:secret@book-postgres:5432/book"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&schema.Book{})

	book := schema.Book{}
	book.Title = "Tris"
	book.Price = 100000
	book.Discount = 20
	book.Rating = 5
	book.Description = "Journal Tris"

	err = db.Create(&book).Error
	if err != nil {
		fmt.Println("==========================")
		fmt.Println("Error Creating Book Record")
		fmt.Println("==========================")
	}

	// fmt.Println("Database Connected")

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/books/:id/:title", handler.GetBooksHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/books", handler.PostBooksHandler)

	router.Run()
}
