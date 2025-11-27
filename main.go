package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Debt struct {
	gorm.Model
	CreditorID   int    `json:"creditorid"`
	Creditor     User	`gorm:"foreignKey:CreditorID"`
	DebtorID     int    `json:"debtorid"`
	Debtor		 User	`gorm:"foreignKey:DebtorID"`
	Amount       int    `json:"amount"`
}

type ShowDebt struct {
	CreditorName string `json:"creditorname"`
	DebtorName   string `json:"debtorname"`
	Amount       int    `json:"amount"`
}

type User struct {
	gorm.Model
	UserName string `json:"username"`
	Password string `json:"password"`
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

// func GetDebts(c *fiber.Ctx) error {
// 	var debts []ShowDebt
// 	for _, debt := range dummyDataBase {
// 		debts = append(debts, ShowDebt{
// 			CreditorName: debt.CreditorName,
// 			DebtorName: debt.DebtorName,
// 			Amount: debt.Amount,
// 		})
// 	}
// 	return c.JSON(debts)
// }

func getStringEnv(key string, fallback string) string {
	if value, exist := os.LookupEnv(key); exist {
		return value
	}
	return fallback
}

func getIntEnv(key string, fallback int) int {
	if value, exist := os.LookupEnv(key); exist {
		if returnValue, err := strconv.Atoi(value); err == nil {
			return returnValue
		}
	}
	return fallback
}

// var (
// 	host = getStringEnv("DB_Host", "localhost")
// 	port = getIntEnv("DB_Port", 5432)
// 	user = getStringEnv("DB_UserName", "postgres")
// 	password = getStringEnv("DB_Password", "sharkyXD")
// 	dbname = "mydatabase"
// )

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Load ENV fail")
	}
	host := getStringEnv("DB_Host", "localhost")
	port := getIntEnv("DB_Port", 5432)
	user := getStringEnv("DB_UserName", "")
	password := getStringEnv("DB_Password", "")
	dbname := "mydatabase"
	fmt.Println("port", port)
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Database connect failed")
	}

	err = db.AutoMigrate(&User{}, &Debt{})
	if err != nil {
		panic(err)
	}

	// app := fiber.New()

	// app.Post("/AddDebt", AddDebt)

	// app.Get("/GetDebts", GetDebts)

	// app.Listen(":8080")
}