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

	bookRepo := aa.NewRepo(db)
	bookPostgres := aa.NewPostgres(bookRepo)

	bookInput := schema.BookInput{
		Title: "dwi",
		Price: "20000",
	}

	bookPostgres.Create(bookInput)

	// books, err := bookRepo.FindAll()
	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("Error Find All Book Record")
	// 	fmt.Println("==========================")
	// }

	// for _, book := range books {
	// 	fmt.Println("Title :", book.Title)
	// }

	// book := schema.Book{}

	// // ================================== create ======================================
	// book.Title = "Tris"
	// book.Price = 100000
	// book.Discount = 20
	// book.Rating = 5
	// book.Description = "Journal Tris"

	// err = db.Create(&book).Error
	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("Error Creating Book Record")
	// 	fmt.Println("==========================")
	// }

	// fmt.Println("Database Connected")

	// // ================================== get first ======================================
	// // var book schema.Book
	// err = db.First(&book).Error
	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("Error Finding Book Record")
	// 	fmt.Println("==========================")
	// }

	// fmt.Println("Title :", book.Title)
	// fmt.Println("Book object %v", book)

	// // =================================== update ========================================
	// err = db.Debug().Where("id = ?", 1).First(&book).Error
	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("Error finding Book Record")
	// 	fmt.Println("==========================")
	// }

	// book.Title = "Anto"

	// err = db.Save(&book).Error
	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("Error Editing Book Record 1")
	// 	fmt.Println("==========================")
	// }

	// // =================================== delete ========================================
	// err = db.Debug().Where("id = ?", 1).First(&book).Error
	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("Error finding book Record")
	// 	fmt.Println("==========================")
	// }
	// err = db.Delete(&book).Error
	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("Error deleting Book Record")
	// 	fmt.Println("==========================")
	// }

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/books/:id/:title", handler.GetBooksHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/books", handler.PostBooksHandler)

	router.Run()
}
