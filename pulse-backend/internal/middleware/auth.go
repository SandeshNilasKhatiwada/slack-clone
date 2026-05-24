package middleware

import (
	"fmt"

	"github.com/SandeshNilasKhatiwada/slack-clone/internal/auth"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *fiber.Ctx) error {
	// Get the token from the request header
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": "Unauthorized",
		})
	}
	tokenString := authHeader[len("Bearer "):]

	// decode the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.NewError(fiber.StatusUnauthorized, "Unexpected signing method")
		}
		return []byte(auth.SecretKey), nil
	})
	fmt.Println(token)
	if err != nil || !token.Valid {
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": "Unauthorized",
		})
	}

	// Move to the next function (the actual route handler)
	return c.Next()
}
