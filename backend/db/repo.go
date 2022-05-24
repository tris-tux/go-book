package aa

import (
	"github.com/tris-tux/go-book/backend/schema"
	"gorm.io/gorm"
)

type Repo interface {
	FindAll() ([]schema.Book, error)
	FindByID(ID int) (schema.Book, error)
	Create(book schema.Book) (schema.Book, error)
}

type repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *repo {
	return &repo{db}
}

func (r *repo) FindAll() ([]schema.Book, error) {
	var books []schema.Book

	err := r.db.Find(&books).Error

	return books, err
}

func (r *repo) FindByID(ID int) (schema.Book, error) {
	var book schema.Book

	err := r.db.Find(&book, ID).Error

	return book, err
}

func (r *repo) Create(book schema.Book) (schema.Book, error) {
	err := r.db.Create(&book).Error

	return book, err
}
