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

var dummyDataBase []Debt

func AddDebt(c *fiber.Ctx) error {
	newDebt := new(Debt)

	if err := c.BodyParser(newDebt); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	dummyDataBase = append(dummyDataBase, *newDebt)

	return c.JSON(newDebt)
}

func main() {
	app := fiber.New()

	app.Post("/AddDebt", AddDebt)

	app.Listen(":8080")
}

// git commit test