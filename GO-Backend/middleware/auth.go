package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// GetUserFromToken extracts user information from JWT token
func GetUserFromToken(c *fiber.Ctx) (map[string]interface{}, error) {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return claims, nil
}

// GetUserID extracts user ID from JWT token
func GetUserID(c *fiber.Ctx) (string, error) {
	claims, err := GetUserFromToken(c)
	if err != nil {
		return "", err
	}
	userID, ok := claims["user_id"].(string)
	if !ok {
		return "", fiber.NewError(fiber.StatusUnauthorized, "Invalid token")
	}
	return userID, nil
}
