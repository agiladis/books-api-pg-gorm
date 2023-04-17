package controllers

import (
	"books-api-pg-gorm/entity"
	"books-api-pg-gorm/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	BookService *services.BookService
}

type BookCreateDto struct {
	NameBook string `json:"name_book" binding:"required"`
	Author   string `json:"author" binding:"required"`
}

func (bc *BookController) CreateBook(ctx *gin.Context) {
	var (
		bookCreateDto BookCreateDto = BookCreateDto{}
		result        entity.Book
		err           error
	)

	if err := ctx.ShouldBindJSON(&bookCreateDto); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	book := entity.Book{
		NameBook: bookCreateDto.NameBook,
		Author:   bookCreateDto.Author,
	}

	if result, err = bc.BookService.Create(book); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": result,
	})

}

func (bc *BookController) GetAllBook(ctx *gin.Context) {
	var (
		books = []entity.Book{}
		err   error
	)

	books, err = bc.BookService.GetAll()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "data not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": books,
	})
}
