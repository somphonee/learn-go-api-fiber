package main

import (
 "github.com/gofiber/fiber/v2"
 "strconv"
)

// Handler functions
// getBooks godoc
// @Summary Get all books
// @Description Get details of all books
// @Tags books
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {array} Book
// @Router /books [get]
func getBooks(c *fiber.Ctx) error {
	return c.JSON(books)
}

func getBook(c *fiber.Ctx) error {

	bookId,err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString(err.Error())
	}
	for _,book := range books{
		if book.ID == bookId {
			return c.JSON(book)
		}
	}
	return c.SendStatus(fiber.StatusNotFound)
}

// Handler functions
// createBooks godoc
// @Summary create books
// @Description Create books
// @Tags books
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 201 {array} Book
// @Router /books [post]
func createBook(c *fiber.Ctx) error {
	book := new(Book)
	if err := c.BodyParser(book); err !=nil{
		return c.Status(fiber.ErrBadRequest.Code).SendString(err.Error())
	}
	books = append(books,*book)

	return c.JSON(books)
} 
// Handler functions
// createBooks godoc
// @Summary edite book bu id
// @Description edite book bu id
// @Tags books
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {array} Book
// @Router /books [put]
func updateBook(c *fiber.Ctx) error {
	bookId,err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString(err.Error())
	}
	bookUpdate := new(Book)
	if err := c.BodyParser(bookUpdate); err !=nil{
		return c.Status(fiber.ErrBadRequest.Code).SendString(err.Error())
	}
	for i,book := range books{
		if book.ID == bookId {
			books[i].Title = bookUpdate.Title
			books[i].Author = bookUpdate.Author
			return c.JSON(books[i])
		}
	}
	return c.SendStatus(fiber.StatusNotFound)
}

// Handler functions
// createBooks godoc
// @Summary delete book bu id
// @Description delete book bu id
// @Tags books
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {array} Book
// @Router /books [delete]
func deleteBook(c *fiber.Ctx) error {
	bookId,err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString(err.Error())
	}
	for i,book := range books{
		if book.ID == bookId {
			books = append(books[:i], books[i+1:]...)
			return c.SendStatus(fiber.StatusNoContent)
		}
	}


	return c.SendStatus(fiber.StatusNotFound)
}