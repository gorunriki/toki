package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("SUPER_SECRET_KEY")

func Protected() fiber.Handler {

	return func(c *fiber.Ctx) error {

		authHeader := c.Get("Authorization")

		if authHeader == "" {
			return c.Status(401).JSON(fiber.Map{
				"error": "missing token",
			})
		}

		tokenString := strings.Replace(
			authHeader,
			"Bearer ",
			"",
			1,
		)

		token, err := jwt.Parse(
			tokenString,
			func(token *jwt.Token) (interface{}, error) {
				return jwtSecret, nil
			},
		)

		if err != nil || !token.Valid {
			return c.Status(401).JSON(fiber.Map{
				"error": "invalid token",
			})
		}

		claims := token.Claims.(jwt.MapClaims)

		c.Locals("user_id", claims["user_id"])
		c.Locals("role", claims["role"])

		return c.Next()
	}
}
