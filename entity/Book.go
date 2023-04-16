package entity

import (
	"errors"

	"gorm.io/gorm"
)

// name convention using singular, will be plural when migrate
type Book struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"not null;type:varchar(50)"`
	Author      string `gorm:"not null;type:varchar(50)"`
	Description string `gorm:"not null;type:text"`
}

// HOOK
func (b *Book) BeforeCreate(tx *gorm.DB) error {
	if b.Title == "" {
		return errors.New("title must be filled")
	}

	if b.Author == "" {
		return errors.New("author must be filled")
	}

	return nil
}
