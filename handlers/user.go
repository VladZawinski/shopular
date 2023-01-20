package handlers

import (
	"shopular/services"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(s *services.UserService) UserHandler {
	return UserHandler{service: s}
}

func (uh UserHandler) FindAllUser(c *fiber.Ctx) error {
	users, err := uh.service.FindAll()
	if err != nil {
		return fiber.NewError(fiber.ErrInternalServerError.Code)
	}
	return c.JSON(users)
}
