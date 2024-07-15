package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/jwt/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

// Book struct to hold book dsata
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}
var books []Book // Slice to store books

func checkMiddleware(c *fiber.Ctx) error {
     start := time.Now()
	fmt.Printf("Request URL: %s - Method: %s - Duration: %s\n", c.OriginalURL(), c.Method(),start)
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	if claims["role"] !="admin" {
		return fiber.ErrUnauthorized
	}
	return c.Next()
}

func main() {
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}
engine := html.New("./views", ".html")
app := fiber.New(fiber.Config{
    Views: engine,
  })


books = append(books, Book{ID: 1,Title: "SPDev",Author: "sp"})
books = append(books, Book{ID: 2,Title: "SP",Author: "sp"})

app.Post("/login", login)

app.Use(checkMiddleware)
// JWT Middleware
app.Use(jwtware.New(jwtware.Config{
	SigningKey: []byte(os.Getenv("JWT_SECRET_KEY")),
}))
app.Use(checkMiddleware)

app.Get("/books",getBooks)
app.Get("/books/:id",getBook)
app.Post("/books",createBook)
app.Put("books/:id",updateBook)
app.Delete("books/:id",deleteBook)
app.Post("/upload",uploadFile)
app.Get("/renderTemplate", renderTemplate)
app.Get("/getConfig", getConfig)
app.Listen(":8080")

}

type User = struct {
	  Email    string `json:"email"`
      Password string `json:"password"`
  }


var memberUser = User {
	Email:    "user@example.com",
	Password: "password123",
  }
  
  func login(c *fiber.Ctx) error {
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if user.Email != memberUser.Email || user.Password != memberUser.Password {
		return fiber.ErrUnauthorized
	  }
	 
	// Create the Claims
	claims := jwt.MapClaims{
		"email": user.Email,
		"role": "admin",
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
  }

