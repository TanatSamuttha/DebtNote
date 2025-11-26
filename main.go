package main

import (
	"github.com/gofiber/fiber/v2"
)

type Debt struct {
	ID           int    `json:"id"`
	CreditorID   int    `json:"creditorid"`
	CreditorName string `json:"creditorname"`
	DebtorID     int    `json:"debtorid"`
	DebtorName   string `json:"debtorname"`
	Amount       int    `json:"amount"`
}

type ShowDebt struct {
	CreditorName string `json:"creditorname"`
	DebtorName   string `json:"debtorname"`
	Amount       int    `json:"amount"`
}

var dummyDataBase []Debt

func AddDebt(c *fiber.Ctx) error {
	newDebt := new(Debt)

	if err := c.BodyParser(newDebt); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	dummyDataBase = append(dummyDataBase, *newDebt)

	return c.JSON(newDebt)
}

func GetDebts(c *fiber.Ctx) error {
	var debts []ShowDebt
	for _, debt := range dummyDataBase {
		debts = append(debts, ShowDebt{
			CreditorName: debt.CreditorName,
			DebtorName: debt.DebtorName,
			Amount: debt.Amount,
		})
	}
	return c.JSON(debts)
}

func main() {
	app := fiber.New()

	app.Post("/AddDebt", AddDebt)

	app.Get("/GetDebts", GetDebts)

	app.Listen(":8080")
}