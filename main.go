package main

import (
	"main.go/core/services"
	"main.go/core/models"

	"fmt"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

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

func getDB_Config(config *models.DB_Config){
	config.Host = getStringEnv("DB_Host", "localhost")
	config.Port = getIntEnv("DB_Port", 5432)
	config.User = getStringEnv("DB_UserName", "")
	config.Password = getStringEnv("DB_Password", "")
	config.DBname = "mydatabase"
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Load ENV fail")
	}
	var db_config models.DB_Config
	getDB_Config(&db_config)
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", db_config.Host, db_config.Port, db_config.User, db_config.Password, db_config.DBname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Database connect failed")
	}

	err = db.AutoMigrate(&models.User{}, &models.Debt{})
	if err != nil {
		panic(err)
	}

	app := fiber.New()

	app.Post("/AddDebt", func (c *fiber.Ctx) error {
		return services.AddDebt(db, c)
	})

	app.Get("/GetDebts", func (c *fiber.Ctx) error  {
		return services.GetDebts(db, c)
	})

	app.Post("/Register", func (c *fiber.Ctx) error {
		return services.Register(db, c)
	})

	app.Listen(":8080")
}