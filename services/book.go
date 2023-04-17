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
