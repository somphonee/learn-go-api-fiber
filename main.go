package main

import (
 "github.com/gofiber/fiber/v2"
)

// Book struct to hold book data
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}
var books []Book // Slice to store books
func main() {
app := fiber.New()
books = append(books, Book{ID: 1,Title: "SPDev",Author: "sp"})
books = append(books, Book{ID: 2,Title: "SP",Author: "sp"})

app.Get("/books",getBooks)
app.Get("/books/:id",getBook)
app.Post("/books",createBook)
app.Put("books/:id",updateBook)
app.Delete("books/:id",deleteBook)
app.Listen(":8080")

}
