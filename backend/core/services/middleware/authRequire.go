package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func getJwtSecretKey(key string, fallback string) string {
	if value, exist := os.LookupEnv(key); exist {
		return value
	}
	return fallback
}

func AuthRequire(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.Parse(cookie, func (token *jwt.Token) (interface{}, error) {
		return []byte(getJwtSecretKey("JWT_SecretKey", "secret")), nil
	})

	claims := token.Claims.(jwt.MapClaims)
	user_id := uint(claims["user_id"].(float64))

	if err != nil || !token.Valid {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	c.Locals("user_id", user_id)

	return c.Next()
}