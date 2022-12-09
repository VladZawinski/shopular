package handlers

import (
	"shopular/dto"
	"shopular/services"

	"github.com/gofiber/fiber/v2"
)

type CategoryHandler struct {
	service *services.CategoryService
}

func NewCategoryHandler(service *services.CategoryService) CategoryHandler {
	return CategoryHandler{
		service: service,
	}
}

func (h *CategoryHandler) GetAll(c *fiber.Ctx) error {
	result, err := h.service.GetAll()
	if err != nil {
		return fiber.NewError(fiber.ErrBadGateway.Code)
	}
	return c.JSON(result)
}

func (h *CategoryHandler) Create(c *fiber.Ctx) error {
	body := new(dto.CreateCategory)
	if err := c.BodyParser(body); err != nil {
		return err
	}
	return c.JSON(h.service.Create(body))
}
