package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func RoleGuard(expected string) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		role := extractUserRole(c)
		if role != expected {
			return fiber.NewError(fiber.StatusUnauthorized)
		}
		return c.Next()
	}
}

func extractUserRole(c *fiber.Ctx) string {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return claims["role"].(string)
}

const (
	AdminRole = "admin"
	UserRole  = "user"
)
