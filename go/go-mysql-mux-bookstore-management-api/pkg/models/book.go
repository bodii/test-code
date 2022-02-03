package models

import (
	"go-mysql-mux-bookstore-management-api/pkg/config"

	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	Id          int64  `json:"bookId"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.Create(b)
	return b
}

func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)

	return books
}

func GetBookById(id int64) (*Book, *gorm.DB) {
	var book Book
	d := db.Where("id=?", id).First(&book)
	return &book, d
}

func DeleteBook(id int64) *Book {
	var book Book
	d := db.Where("id=?", id).First(&book)
	if book.Id == 0 {
		return &book
	}

	d.Delete(&Book{})
	return &book
}
