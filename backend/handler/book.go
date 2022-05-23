package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/tris-tux/go-book/backend/schema"
)

func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Trisno",
		"bio":  "Software Engineer",
	})
}

func GetBooksHandler(c *gin.Context) {
	id := c.Param("id")
	title := c.Param("title")

	c.JSON(http.StatusOK, gin.H{"id": id, "title": title})
}

func QueryHandler(c *gin.Context) {
	id := c.Query("id")

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func PostBooksHandler(c *gin.Context) {
	var bookInput schema.BookInput

	err := c.ShouldBindJSON(&bookInput)
	if err != nil {

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on Field %s, conditional: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})

		// c.JSON(http.StatusBadRequest, http.StatusBadRequest)
		return
		// log.Fatal(err) // fatal itu langsung mati
	}

	c.JSON(http.StatusOK, gin.H{
		"title": bookInput.Title,
		"price": bookInput.Price,
	})
}
