package handlers

import (
	"shopular/services"

	"github.com/gofiber/fiber/v2"
)

type CustomerHandler struct {
	service *services.CustomerService
}

func NewCustomerHandler(s *services.CustomerService) CustomerHandler {
	return CustomerHandler{service: s}
}

func (h CustomerHandler) FindAll(c *fiber.Ctx) error {
	result, err := h.service.FindAll()
	if err != nil {
		return fiber.NewError(fiber.ErrBadGateway.Code)
	}
	return c.JSON(result)
}
