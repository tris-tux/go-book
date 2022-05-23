package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tris-tux/go-book/backend/handler"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/books/:id/:title", handler.GetBooksHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/books", handler.PostBooksHandler)

	router.Run()
}
