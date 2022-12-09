package handlers

import (
	"shopular/dto"
	"shopular/services"

	"github.com/gofiber/fiber/v2"
)

type SubCategoryHandler struct {
	service *services.SubCategoryService
}

func NewSubCategoryHandler(service *services.SubCategoryService) SubCategoryHandler {
	return SubCategoryHandler{
		service: service,
	}
}

func (h *SubCategoryHandler) GetAll(c *fiber.Ctx) error {
	result, err := h.service.GetAll()
	if err != nil {
		return fiber.NewError(fiber.ErrBadGateway.Code)
	}
	return c.JSON(result)
}

func (h *SubCategoryHandler) Create(c *fiber.Ctx) error {
	body := new(dto.CreateSubCategory)
	if err := c.BodyParser(body); err != nil {
		return err
	}
	return c.JSON(h.service.Create(body))
}
