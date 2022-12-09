package handlers

import (
	"shopular/dto"
	"shopular/services"

	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	service *services.ProductService
}

func NewProductHandler(service *services.ProductService) ProductHandler {
	return ProductHandler{
		service: service,
	}
}

func (h *ProductHandler) GetAll(c *fiber.Ctx) error {
	result, err := h.service.GetAll()
	if err != nil {
		return fiber.NewError(fiber.ErrBadGateway.Code)
	}
	return c.JSON(result)
}

func (h *ProductHandler) Create(c *fiber.Ctx) error {
	body := new(dto.CreateProduct)
	if err := c.BodyParser(body); err != nil {
		return err
	}
	return c.JSON(h.service.Create(body))
}
