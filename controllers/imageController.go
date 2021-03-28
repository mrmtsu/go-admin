package controllers

import (
	"go-admin/models"
	"os"

	"github.com/gofiber/fiber/v2"
)

func Upload(c *fiber.Ctx) error {
	models.Env_load()

	form, err := c.MultipartForm()

	if err != nil {
		return err
	}

	files := form.File["image"]
	filename := ""

	for _, file := range files {
		filename = file.Filename

		if err := c.SaveFile(file, "./uploads/"+filename); err != nil {
			return err
		}
	}

	return c.JSON(fiber.Map{
		"url": os.Getenv("HTTP_UPLOAD") + filename,
	})
}
