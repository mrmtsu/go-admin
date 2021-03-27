package controllers

import (
	"go-admin/database"
	"go-admin/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func AllOrders(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))

	return c.JSON(models.Paginate(database.DB, &models.Order{}, page))
}
