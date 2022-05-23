package main

import (
	"./handler"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/books/:id/:title", handler.GetBooksHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/books", hanler.PostBooksHandler)

	router.Run()
}