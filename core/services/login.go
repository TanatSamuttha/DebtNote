package services

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"main.go/core/models"
)

func Login(db *gorm.DB, c *fiber.Ctx) error {
	var loginUser models.User
	var dbUser models.User
	
	if err := c.BodyParser(&loginUser); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	result := db.Table("users").Where("user_name = ?", loginUser.UserName).First(&dbUser)
	if result.Error != nil {
		return c.SendString("User or password is incorrect")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(loginUser.Password)); err != nil {
		return c.SendString("User or password is incorrect")
	}

	return c.SendStatus(fiber.StatusOK)
}