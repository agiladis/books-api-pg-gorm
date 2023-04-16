package services

import "gorm.io/gorm"

type BookService struct {
	DB *gorm.DB
}
