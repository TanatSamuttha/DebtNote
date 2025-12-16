package services

import (
	"debtnote/core/models"
	
	"gorm.io/gorm"
	"github.com/gofiber/fiber/v2"
)

func MyDebts(db *gorm.DB, c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	var debts []models.ShowDebt
	db.Table("debts").
		Select(`debts.id, debts.amount, creditors.user_name AS creditor_name, debtors.user_name AS debtor_name`).
		Joins("JOIN users AS creditors ON creditors.id = debts.creditor_id").
		Joins("JOIN users AS debtors ON debtors.id = debts.debtor_id").
		Where("debts.creditor_id = ? OR debts.debtor_id = ?", userID, userID).
		Scan(&debts)
	return c.JSON(debts)
}