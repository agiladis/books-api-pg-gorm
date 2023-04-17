package services

import (
	"books-api-pg-gorm/entity"

	"gorm.io/gorm"
)

type BookService struct {
	DB *gorm.DB
}

func (bs *BookService) Create(book entity.Book) (entity.Book, error) {
	if err := bs.DB.Create(&book).Error; err != nil {
		return entity.Book{}, err
	}

	return book, nil
}

func (bs *BookService) GetAll() ([]entity.Book, error) {
	var books []entity.Book

	if err := bs.DB.Find(&books).Error; err != nil {
		return books, err
	}

	return books, nil
}
