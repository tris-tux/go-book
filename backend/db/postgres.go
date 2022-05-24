package aa

import "github.com/tris-tux/go-book/backend/schema"

type Postgres interface {
	FindAll() ([]schema.Book, error)
	FindByID(ID int) (schema.Book, error)
	Create(bookInput schema.BookInput) (schema.Book, error)
	Update(ID int, bookInput schema.BookInput) (schema.Book, error)
	Delete(ID int) (schema.Book, error)
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
}

func (p *postgres) FindByID(ID int) (schema.Book, error) {
	book, err := p.repo.FindByID(ID)
	return book, err
}

func (p *postgres) Create(bookInput schema.BookInput) (schema.Book, error) {
	price, _ := bookInput.Price.Int64()
	rating, _ := bookInput.Rating.Int64()
	discount, _ := bookInput.Discount.Int64()

	book := schema.Book{
		Title:       bookInput.Title,
		Price:       int(price),
		Description: bookInput.Description,
		Rating:      int(rating),
		Discount:    int(discount),
	}

	newBook, err := p.repo.Create(book)
	return newBook, err
}

func (p *postgres) Update(ID int, bookInput schema.BookInput) (schema.Book, error) {
	book, err := p.repo.FindByID(ID)
	price, _ := bookInput.Price.Int64()
	rating, _ := bookInput.Rating.Int64()
	discount, _ := bookInput.Discount.Int64()

	book.Title = bookInput.Title
	book.Price = int(price)
	book.Description = bookInput.Description
	book.Rating = int(rating)
	book.Discount = int(discount)

	newBook, err := p.repo.Update(book)
	return newBook, err
}

func (p *postgres) Delete(ID int) (schema.Book, error) {
	book, err := p.repo.FindByID(ID)

	newBook, err := p.repo.Delete(book)
	return newBook, err
}
