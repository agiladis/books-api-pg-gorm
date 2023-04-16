package main

import (
	"books-api-pg-gorm/entity"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	DB_USERNAME = "postgres"
	DB_PASSWORD = "postgres"
	DB_HOST     = "localhost"
	DB_PORT     = "5432"
	DB_DATABASE = "db_book"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	// postgres://postgres:postgres@localhost:5432/db_book?sslmode=disable
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", DB_USERNAME, DB_PASSWORD, DB_HOST, DB_PORT, DB_DATABASE)

	db, err = gorm.Open(postgres.Open(connString), &gorm.Config{
		// Log to see changes
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  logger.Silent,
				IgnoreRecordNotFoundError: true,
				Colorful:                  true,
			},
		),
	})

	if err != nil {
		panic(err)
	}
}

func main() {

	db.Debug().AutoMigrate(&entity.Book{})

}
