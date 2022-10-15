package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"errors"
)

type Book struct {
	ID       string `json:"id"`
	Title	 string `json:"title"`
	Author	 string `json:"author"`
	Stock	 int 	`json:"stock"`
}

var books = []Book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Stock: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Stock: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Stock: 6},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func createBooks(c *gin.Context) {
	var newBooks Book

	if err := c.BindJSON(&newBooks); err != nil {
		return
	}

	books = append(books, newBooks)
	c.IndentedJSON(http.StatusOK, books)
}

func bookById(c *gin.Context) {
	id := c.Param("id")
	book, err := getBooksById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func getBooksById(id string) (Book, error) {
	for i, b := range books {
		if b.ID == id {
			return books[i], nil
		}
	}

	return nil, errors.New("Book not found")
}

func main() {
	r := gin.Default()

	r.GET("/books", getBooks)
	r.POST("/create_books", createBooks)
	r.GET("/books/:id", bookById)

	r.Run()
}