package services

import (
	"main.go/core/models"

	"gorm.io/gorm"
	"github.com/gofiber/fiber/v2"
)

func getUserID(db *gorm.DB, username string) uint {
	var user models.User
	db.Table("users").Where("user_name = ?", username).First(&user)
	return user.ID
}

func AddDebt(db *gorm.DB, c *fiber.Ctx) error {
	inputDebt := new(models.ShowDebt)

	if err := c.BodyParser(inputDebt); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	newDebt := models.Debt{CreditorID: getUserID(db, inputDebt.CreditorName),
		DebtorID: getUserID(db, inputDebt.DebtorName),
		Amount: inputDebt.Amount,
	}

	result := db.Create(&newDebt)
	if result.Error != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.SendStatus(fiber.StatusOK)
}