package services

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"main.go/core/models"
)

func usernameIsExist(db *gorm.DB, username string) (bool, error) {
	var count int64
	result := db.Table("users").Where("user_name = ?", username).Count(&count)
	return (count > 0), result.Error
}

func Register(db *gorm.DB, c *fiber.Ctx) error{
	newUser := new(models.User)
	if err := c.BodyParser(newUser); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	
	isExist, err := usernameIsExist(db, newUser.UserName);
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	} else if  isExist {
		return c.SendString("Registor fail: Username is already exist")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	newUser.Password = string(hashedPassword)
	result := db.Create(newUser)
	if result.Error != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusOK)
}