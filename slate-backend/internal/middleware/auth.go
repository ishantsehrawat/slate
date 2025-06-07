package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/ishant/slate-backend/config"
)

func RequireAuth(c *fiber.Ctx) error {
	// ✅ Get token from cookie instead of header
	tokenStr := c.Cookies("auth_token")
	if tokenStr == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Missing auth token",
		})
	}

	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		// Validate the alg is what we expect
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.NewError(fiber.StatusUnauthorized, "Unexpected signing method")
		}
		return config.JWTSecret, nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid or expired token",
		})
	}

	// ✅ Store useful claims in context
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		c.Locals("user_id", claims["user_id"])
		c.Locals("email", claims["email"])
	}

	return c.Next()
}
