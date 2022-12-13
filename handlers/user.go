package handlers

import (
	"shopular/services"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	service services.UserService
}

func (uh UserHandler) FindAllUser(c fiber.Ctx) {

}
