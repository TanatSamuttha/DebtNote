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
	CreditorID   uint    `json:"creditorid"`
	Creditor     User	 `gorm:"foreignKey:CreditorID"`
	DebtorID     uint    `json:"debtorid"`
	Debtor		 User	 `gorm:"foreignKey:DebtorID"`
	Amount       uint    `json:"amount"`
}

type ShowDebt struct {
	gorm.Model
	CreditorName string `json:"creditorname"`
	DebtorName   string `json:"debtorname"`
	Amount       uint   `json:"amount"`
}

type User struct {
	gorm.Model
	UserName string `json:"username"`
	Password string `json:"password"`
}

var dummyDataBase []Debt

func AddDebt(db *gorm.DB, c *fiber.Ctx) error {
	inputDebt := new(ShowDebt)

	if err := c.BodyParser(inputDebt); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	newDebt := Debt{CreditorID: getUserID(db, inputDebt.CreditorName),
		DebtorID: getUserID(db, inputDebt.DebtorName),
		Amount: inputDebt.Amount,
	}

	result := db.Create(&newDebt)
	if result.Error != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.SendStatus(fiber.StatusOK)
}

func getUserID(db *gorm.DB, username string) uint {
	var user User
	db.Where("user_name = ?", username).First(&user)
	return user.ID
}

func GetDebts(db *gorm.DB, c *fiber.Ctx) error {
	var debts []ShowDebt
	db.Find(&debts)
	return c.JSON(debts)
}

func getStringEnv(key string, fallback string) string {
	if value, exist := os.LookupEnv(key); exist {
		return value
	}
	return fallback
}

func AddUser(db *gorm.DB, c *fiber.Ctx) error {
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	result := db.Create(&user)
	if result.Error != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.SendStatus(fiber.StatusOK)
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
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Database connect failed")
	}

	err = db.AutoMigrate(&User{}, &Debt{})
	if err != nil {
		panic(err)
	}

	app := fiber.New()

	app.Post("/AddDebt", func (c *fiber.Ctx) error {
		return AddDebt(db, c)
	})

	app.Get("/GetDebts", func (c *fiber.Ctx) error  {
		return GetDebts(db, c)
	})

	app.Post("/AddUser", func (c *fiber.Ctx) error {
		return AddUser(db, c)
	})

	app.Listen(":8080")
}