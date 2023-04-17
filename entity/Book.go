package entity

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

// name convention using singular, will be plural when migrate
type Book struct {
	ID        uint   `gorm:"primaryKey"`
	NameBook  string `gorm:"not null;type:varchar(50)"`
	Author    string `gorm:"not null;type:varchar(50)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// HOOK
func (b *Book) BeforeCreate(tx *gorm.DB) error {
	if b.NameBook == "" {
		return errors.New("book name must be filled")
	}

	if b.Author == "" {
		return errors.New("author must be filled")
	}

	return nil
}
