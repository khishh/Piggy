package repository

import (
	"github.com/khishh/personal-finance-app/app/models"
	"github.com/khishh/personal-finance-app/graph/model"
	"gorm.io/gorm"
)

type BookRepository interface {
	CreateBook(bookInput *model.BookInput) (*models.Book, error)
	UpdateBook(bookInput *model.BookInput, id int) error
	DeleteBook(id int) error
	GetOneBook(id int) (*models.Book, error)
	GetAllBooks() ([]*model.Book, error)
}

type BookService struct {
	Db *gorm.DB
}

var _ BookRepository = &BookService{}

func NewBookService(db *gorm.DB) *BookService {
	return &BookService{
		Db: db,
	}
}

func (b *BookService) CreateBook(bookInput *model.BookInput) (*models.Book, error) {
	book := &models.Book{
		Title:     bookInput.Title,
		Author:    bookInput.Author,
		Publisher: bookInput.Publisher,
	}
	err := b.Db.Create(&book).Error

	return book, err
}

func (b *BookService) UpdateBook(bookInput *model.BookInput, id int) error {
	book := models.Book{
		ID:        id,
		Title:     bookInput.Title,
		Author:    bookInput.Author,
		Publisher: bookInput.Publisher,
	}
	err := b.Db.Model(&book).Where("id = ?", id).Updates(book).Error
	return err
}

func (b *BookService) DeleteBook(id int) error {
	book := &models.Book{}
	err := b.Db.Delete(book, id).Error
	return err
}

func (b *BookService) GetOneBook(id int) (*models.Book, error) {
	book := &models.Book{}
	// change b.Db.First(&models.Book{}, id).Error?
	err := b.Db.Where("id = ?", id).First(book).Error
	return book, err
}

// change model.Book --> models.Book?
func (b *BookService) GetAllBooks() ([]*model.Book, error) {
	books := []*model.Book{}
	err := b.Db.Find(&books).Error
	return books, err

}
