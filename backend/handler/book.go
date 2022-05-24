package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	aa "github.com/tris-tux/go-book/backend/db"
	"github.com/tris-tux/go-book/backend/schema"
)

type bookHandler struct {
	bookPostgres aa.Postgres
}

func NewBookHandler(bookPostgres aa.Postgres) *bookHandler {
	return &bookHandler{bookPostgres}
}

func (h *bookHandler) RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Trisno",
		"bio":  "Software Engineer",
	})
}

func (h *bookHandler) GetBooksHandler(c *gin.Context) {
	id := c.Param("id")
	title := c.Param("title")

	c.JSON(http.StatusOK, gin.H{"id": id, "title": title})
}

func (h *bookHandler) QueryHandler(c *gin.Context) {
	id := c.Query("id")

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (h *bookHandler) PostBooksHandler(c *gin.Context) {
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

	book, err := h.bookPostgres.Create(bookInput)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}
