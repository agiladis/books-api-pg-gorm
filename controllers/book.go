package controllers

import "books-api-pg-gorm/services"

type BookController struct {
	BookService *services.BookService
}

type BookCreateDto struct {
	NameBook string `json:"name_book" binding:"required"`
	Author   string `json:"author" binding:"required"`
}
