package services

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"main.go/core/models"
)

func Register(db *gorm.DB, c *fiber.Ctx) error{
	newUser := new(models.User)
	if err := c.BodyParser(newUser); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	newUser.Password = string(hashedPassword)
	result := db.Create(newUser)
	if result.Error != nil {
		c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusOK)
}