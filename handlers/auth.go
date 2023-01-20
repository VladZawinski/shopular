package handlers

import (
	"fmt"
	"shopular/dto"
	"shopular/services"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type AuthHandler struct {
	userService     services.UserService
	customerService services.CustomerService
}

func NewAuthHandler(service services.UserService) AuthHandler {
	return AuthHandler{userService: service}
}

func (h AuthHandler) Login(c *fiber.Ctx) error {
	body := new(dto.Login)
	if err := c.BodyParser(body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Fields are required")
	}
	u := h.userService.FindByUsername(body.Username)
	if u == nil {
		return fiber.NewError(fiber.StatusForbidden)
	}
	if u.Password != body.Password {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid password", "data": nil})
	}
	claims := jwt.MapClaims{
		"username": u.Username,
		"userId":   u.ID,
		"role":     u.Role,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}
	fmt.Println(claims)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.JSON(fiber.Map{"token": t})
}

func (h AuthHandler) Register(c *fiber.Ctx) error {
	body := new(dto.RegisterUser)
	if err := c.BodyParser(body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Fields are required")
	}
	existing := h.userService.FindByUsername(body.Username)
	if existing != nil {
		return fiber.NewError(fiber.StatusConflict)
	}
	user := h.userService.CreateNewUser(*body)
	if user != nil {
		h.customerService.Create(dto.CreateCustomer{
			FirstName: body.FirstName,
			LastName:  body.LastName,
			Address:   body.Address,
			Phone:     body.Phone,
			Email:     body.Email,
		})
	}
	return c.SendStatus(fiber.StatusCreated)
}
