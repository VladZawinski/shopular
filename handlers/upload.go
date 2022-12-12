package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type UploadHandler struct{}

func NewUploadHandler() UploadHandler {
	return UploadHandler{}
}

func (h UploadHandler) UploadProductImage(c *fiber.Ctx) error {
	file, err := c.FormFile("image")
	if err == nil {
		return c.SaveFile(file, fmt.Sprintf("./images/%s", file.Filename))
	}
	return err
}
