package controllers

import "books-api-pg-gorm/services"

type BookController struct {
	BookService *services.BookService
}

type BookCreateDto struct {
	Title       string `json:"title" binding:"required"`
	Author      string `json:"author" binding:"required"`
	Description string `json:"description"`
}
