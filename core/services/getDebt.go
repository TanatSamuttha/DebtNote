package services

import (
	"main.go/core/models"
	
	"gorm.io/gorm"
	"github.com/gofiber/fiber/v2"
)

func GetDebts(db *gorm.DB, c *fiber.Ctx) error {
	var debts []models.ShowDebt
	db.Find(&debts)
	return c.JSON(debts)
}