package router

import (
	"books-api-pg-gorm/controllers"
	"books-api-pg-gorm/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func StartServer(db *gorm.DB) *gin.Engine {
	bookService := services.BookService{
		DB: db,
	}
	// userController will use userService to get data from db
	bookController := controllers.BookController{
		BookService: &bookService,
	}

	app := gin.Default()

	return app
}
