package services

import (
	"time"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"main.go/core/models"
)

func getJwtSecretKey(key string, fallback string) string {
	if value, exist := os.LookupEnv(key); exist {
		return value
	}
	return fallback
}

func Login(db *gorm.DB, c *fiber.Ctx) error {
	var loginUser models.User
	var dbUser models.User
	
	if err := c.BodyParser(&loginUser); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	result := db.Table("users").Where("user_name = ?", loginUser.UserName).First(&dbUser)
	if result.Error != nil {
		return c.SendString("User or password is incorrect")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(loginUser.Password)); err != nil {
		return c.SendString("User or password is incorrect")
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = loginUser.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	tokenValue, err := token.SignedString(getJwtSecretKey("JWT_SecretKey", "secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	c.Cookie(&fiber.Cookie{
		Name:	  "jwt",
		Value:    tokenValue,
		Expires:  time.Now().Add(time.Hour * 72),
		HTTPOnly: true,
	})

	return c.SendStatus(fiber.StatusOK)
}