package middleware

import (
	"github.com/SandeshNilasKhatiwada/slack-clone/internal/service"
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
		return []byte(service.SecretKey), nil
	})
	if err != nil || !token.Valid {
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": "Unauthorized",
		})
	}

	// Extract claims and store in context locals
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		c.Locals("user", claims)
	}

	// Move to the next function (the actual route handler)
	return c.Next()
}
