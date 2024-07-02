package main

import (
 "github.com/gofiber/fiber/v2"
 "os"
)
func getConfig(c *fiber.Ctx) error {
	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		secretKey = "defaultSecret" 
	  }
	return c.JSON(fiber.Map{
		"SECRET_KEY": secretKey,
	  })
  }