package services

import (
	"books-api-pg-gorm/entity"

	"errors"

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

func (bs *BookService) GetBookById(id uint) (entity.Book, error) {
	var book entity.Book

	if err := bs.DB.Where("id = ?", id).First(&book).Error; err != nil {
		return book, err
	}

	return book, nil
}

func (bs *BookService) Update(id uint, book entity.Book) (entity.Book, error) {
	// result := bs.DB.Model(&entity.Book{}).Where("id = ?", id).Updates(&book)
	var newBook entity.Book

	result := bs.DB.Model(&newBook).Where("id = ?", id).Updates(&book)

	if result.RowsAffected == 0 {
		return entity.Book{}, errors.New("there is no data to update")
	}

	return newBook, nil
}

func (bs *BookService) Delete(id uint) error {
	var book entity.Book

	if err := bs.DB.Where("id = ?", id).Delete(&book).Error; err != nil {
		return err
	}

	return nil
}
