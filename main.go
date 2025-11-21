package main

import "github.com/gofiber/fiber"

type Debt struct {
	ID           int    `json: "id"`
	creditorID   int    `json: creditor_id`
	creditorName string `json: "creditor_name"`
	debtorID     int    `json: debtor_id`
	debtorName   string `json: debtor_name`
	amount       int    `json: amount`
}

func main() {
	app := fiber.New()

	app.Listen(":8080")
}