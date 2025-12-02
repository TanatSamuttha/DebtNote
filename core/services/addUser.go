package services

import (
	"main.go/core/models"

	"gorm.io/gorm"
	"github.com/gofiber/fiber/v2"
)

func AddUser(db *gorm.DB, c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	result := db.Create(&user)
	if result.Error != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.SendStatus(fiber.StatusOK)
}