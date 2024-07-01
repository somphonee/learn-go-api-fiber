package main

import (
 "github.com/gofiber/fiber/v2"
 "github.com/gofiber/template/html/v2"
)

// Book struct to hold book data
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}
var books []Book // Slice to store books
func main() {
engine := html.New("./views", ".html")
app := fiber.New(fiber.Config{
    Views: engine,
  })
books = append(books, Book{ID: 1,Title: "SPDev",Author: "sp"})
books = append(books, Book{ID: 2,Title: "SP",Author: "sp"})



app.Get("/books",getBooks)
app.Get("/books/:id",getBook)
app.Post("/books",createBook)
app.Put("books/:id",updateBook)
app.Delete("books/:id",deleteBook)
app.Post("/upload",uploadFile)

app.Get("/renderTemplate", renderTemplate)
app.Listen(":8080")

}

