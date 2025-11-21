package main

import "github.com/gofiber/fiber/v2"

type Debt struct {
	ID           int    `json: "id"`
	creditorID   int    `json: creditor_id`
	creditorName string `json: "creditor_name"`
	debtorID     int    `json: debtor_id`
	debtorName   string `json: debtor_name`
	amount       int    `json: amount`
}

func AddDebt(c *fiber.Ctx) error {
	newDebt := new(Debt)

	if err := c.BodyParser(newDebt); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return c.JSON(newDebt)
}

func main() {
	app := fiber.New()

	app.Listen(":8080")
}