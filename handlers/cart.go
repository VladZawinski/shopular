package handlers

import (
	"shopular/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type CartHandler struct {
	service services.CartService
}

func NewCartHandler(service services.CartService) CartHandler {
	return CartHandler{service: service}
}

func (h CartHandler) ViewCart(c *fiber.Ctx) error {
	userId := ExtractUserId(c)
	carts, err := h.service.FindByUserId(int(userId))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest)
	}
	return c.JSON(carts)
}

func (h CartHandler) AddToCart(c *fiber.Ctx) error {
	userId := ExtractUserId(c)
	productId := c.Params("productId")
	i, err := strconv.Atoi(productId)
	if err != nil {
		panic(err)
	}
	return h.service.AddToCart(int(userId), i)
}

func (h CartHandler) RemoveFromCart(c *fiber.Ctx) error {
	cartId := c.Params("cartId")
	i, err := strconv.Atoi(cartId)
	if err != nil {
		panic(err)
	}
	return h.service.RemoveFromCart(i)
}
