package aa

import "github.com/tris-tux/go-book/backend/schema"

type Postgres interface {
	FindAll() ([]schema.Book, error)
	FindByID(ID int) (schema.Book, error)
	Create(bookInput schema.BookInput) (schema.Book, error)
}

type postgres struct {
	repo Repo
}

func NewPostgres(repo Repo) *postgres {
	return &postgres{repo}
}

func (p *postgres) FindAll() ([]schema.Book, error) {
	books, err := p.repo.FindAll()
	return books, err
	// return p.repo.FindAll()
}

func (p *postgres) FindByID(ID int) (schema.Book, error) {
	book, err := p.repo.FindByID(ID)
	return book, err
	// return p.repo.FindAll()
}

func (p *postgres) Create(bookInput schema.BookInput) (schema.Book, error) {
	price, _ := bookInput.Price.Int64()

	book := schema.Book{
		Title: bookInput.Title,
		Price: int(price),
	}

	newBook, err := p.repo.Create(book)
	return newBook, err
}
