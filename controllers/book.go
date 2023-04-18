package controllers

import (
	"books-api-pg-gorm/entity"
	"books-api-pg-gorm/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	BookService *services.BookService
}

type BookCreateDto struct {
	NameBook string `json:"name_book" binding:"required"`
	Author   string `json:"author" binding:"required"`
}

type BookUpdateDto struct {
	NameBook string `json:"name_book"`
	Author   string `json:"author"`
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

	ctx.JSON(http.StatusCreated, gin.H{
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

func (bc *BookController) GetBookById(ctx *gin.Context) {
	var (
		bookId = ctx.Param("id")
		book   entity.Book
	)

	bookIdInt, _ := strconv.Atoi(bookId)
	book, err := bc.BookService.GetBookById(uint(bookIdInt))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "id not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": book,
	})

}

func (bc *BookController) UpdateBookById(ctx *gin.Context) {
	var (
		bookUpdateDto BookUpdateDto = BookUpdateDto{}
		result        entity.Book
		err           error
	)

	if err := ctx.ShouldBindJSON(&bookUpdateDto); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	bookId := ctx.Param("id")
	bookIdInt, _ := strconv.Atoi(bookId)
	book := entity.Book{
		NameBook: bookUpdateDto.NameBook,
		Author:   bookUpdateDto.Author,
	}

	result, err = bc.BookService.Update(uint(bookIdInt), book)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"data": result,
	})
}

func (bc *BookController) DeleteBookById(ctx *gin.Context) {
	bookId := ctx.Param("id")
	bookIdInt, _ := strconv.Atoi(bookId)

	if err := bc.BookService.Delete(uint(bookIdInt)); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "id not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "book deleted successfully",
	})
}
