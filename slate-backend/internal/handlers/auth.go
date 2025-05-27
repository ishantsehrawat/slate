package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/ishant/slate-backend/config"
	"github.com/ishant/slate-backend/internal/database"
	"github.com/ishant/slate-backend/internal/models"
	"github.com/ishant/slate-backend/internal/services"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
)


func GoogleLogin(c *fiber.Ctx) error {
	// Generate the Google OAuth2 URL to redirect the user to Google login page
	fmt.Println("huwhuw", services.GoogleOauthConfig.ClientID)
	fmt.Print(config.GoogleClientID)
	url := services.GoogleOauthConfig.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Println(url)
	return c.Redirect(url, fiber.StatusTemporaryRedirect)
}

func GoogleCallback(c *fiber.Ctx) error {
	code := c.Query("code")
	if code == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Code not found in query"})
	}

	userData, err := services.ExchangeCodeForUser(context.Background(), code)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get user info from Google"})
	}

	// Check if user exists in DB, else create
	var user models.User
	result := database.DB.Where("google_id = ?", userData.GoogleID).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			// Create new user
			user = *userData
			if err := database.DB.Create(&user).Error; err != nil {
				return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create user"})
			}
		} else {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "DB error"})
		}
	}

	token, err := services.GenerateJWT(&user)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate token"})
	}

	// Return token, usually in JSON or set cookie
	return c.JSON(fiber.Map{"token": token})
}
