package handlers

import (
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/ishant/slate-backend/internal/database"
	"github.com/ishant/slate-backend/internal/models"
	"github.com/ishant/slate-backend/internal/services"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
)

func GoogleLogin(c *fiber.Ctx) error {
	// Generate the OAuth2 URL to redirect user to Google's login
	url := services.GoogleOauthConfig.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	return c.Redirect(url, fiber.StatusTemporaryRedirect)
}

func GoogleCallback(c *fiber.Ctx) error {
	code := c.Query("code")
	if code == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Authorization code not found"})
	}

	userData, err := services.ExchangeCodeForUser(context.Background(), code)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch user info from Google"})
	}

	var user models.User
	result := database.DB.Where("google_id = ?", userData.GoogleID).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			user = *userData
			if err := database.DB.Create(&user).Error; err != nil {
				return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create user in database"})
			}
		} else {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Database error"})
		}
	}

	token, err := services.GenerateJWT(&user)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate JWT token"})
	}

	// Set the JWT in a secure, HttpOnly cookie for authentication
	c.Cookie(&fiber.Cookie{
		Name:     "auth_token",
		Value:    token,
		HTTPOnly: true,
		Secure:   false, // ⚠️ Use true in production with HTTPS!
		Path:     "/",
		SameSite: "Lax",            // Consider Strict or None if cross-site and secure is true
		MaxAge:   60 * 60 * 24 * 7, // 7 days expiry
	})

	// Redirect to frontend after successful login
	return c.Redirect("http://test.slate.com:5173/auth/callback", fiber.StatusTemporaryRedirect)
}
