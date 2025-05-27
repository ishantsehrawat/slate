package utils

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// GetUserIDFromContext safely extracts the user ID as an integer
func GetUserIDFromContext(c *fiber.Ctx) (int, error) {
	userIDRaw := c.Locals("user_id")
	switch v := userIDRaw.(type) {
	case float64:
		return int(v), nil
	case string:
		return strconv.Atoi(v)
	default:
		return 0, fiber.NewError(fiber.StatusUnauthorized, "Invalid user ID in token")
	}
}
