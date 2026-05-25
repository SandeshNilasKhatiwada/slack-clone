package auth

import (
	"fmt"
	"time"

	"github.com/SandeshNilasKhatiwada/slack-clone/internal/database"
	"github.com/SandeshNilasKhatiwada/slack-clone/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

const SecretKey = "your_secret_key_here"

func Register(c *fiber.Ctx) error {

	// temporary interface to hold the request body
	var data map[string]interface{}

	if err := c.BodyParser(&data); err != nil {
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	// password with bcrypt hash
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"].(string)), 14)

	//create the new data with the password hash on it and the rest of the data from the request body
	user := models.User{
		Username: data["username"].(string),
		Email:    data["email"].(string),
		Password: string(password),
	}

	// save the user to the database
	if err := database.GetDB().Create(&user).Error; err != nil {
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to create user",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "User successfully registered",
		"user": fiber.Map{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
		},
	})
}

func Login(c *fiber.Ctx) error {
	var data map[string]interface{}

	if err := c.BodyParser(&data); err != nil {
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}
	var user models.User

	//find user by email
	database.GetDB().Where("email = ?", data["email"].(string)).First(&user)

	// if user not found send error response
	if user.ID == 0 {
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": "User not found",
		})
	}

	// compare password with bcrypt
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"].(string))); err != nil {
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid password",
			"error":   err.Error(),
		})
	}

	// create the claims for the JWT tokens

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(SecretKey))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not login"})
	}

	return c.JSON(fiber.Map{
		"message": "Login successful",
		"user": fiber.Map{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
		},
		"token": token,
	})
}

func UpdateUser(c *fiber.Ctx) error {
	fmt.Println(user)

	return c.JSON(fiber.Map{
		"message": "Update user",
	})
}
