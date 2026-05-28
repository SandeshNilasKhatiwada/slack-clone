package auth

import (
	"github.com/SandeshNilasKhatiwada/slack-clone/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type AuthHandler struct {
	svc service.AuthService
}

func NewAuthHandler(svc service.AuthService) *AuthHandler {
	return &AuthHandler{svc: svc}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	user, err := h.svc.Register(data["username"], data["email"], data["password"])
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"message": "User successfully registered",
		"user":    fiber.Map{"id": user.ID, "username": user.Username, "email": user.Email},
	})
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	token, user, err := h.svc.Login(data["email"], data["password"])
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"message": "Login successful",
		"user":    fiber.Map{"id": user.ID, "username": user.Username, "email": user.Email},
		"token":   token,
	})
}

func (h *AuthHandler) GetMe(c *fiber.Ctx) error {
	claims, ok := c.Locals("user").(jwt.MapClaims)
	if !ok {
		return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
	}

	id := uint(claims["id"].(float64))
	user, err := h.svc.GetUserByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	return c.JSON(fiber.Map{
		"message": "User details retrieved successfully",
		"user":    fiber.Map{"id": user.ID, "username": user.Username, "email": user.Email},
	})
}
