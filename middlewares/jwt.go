package middlewares

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func JwtGuard() func(*fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	})
}
