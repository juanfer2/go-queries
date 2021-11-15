package controllers

import (
	"juanfer2/go-queries/src/modules/queries/services"
	"juanfer2/go-queries/src/types"

	"github.com/gofiber/fiber/v2"
)

func parserBodyQuery(c *fiber.Ctx) types.PayloadQuery {
	payload := types.PayloadQuery{}

	if err := c.BodyParser(&payload); err != nil {
		// return err
	}

	return payload
}

func GetQueries(c *fiber.Ctx) error {
	list := services.GetAllQueries()
	return c.JSON(list)
}

func PostCreateQuery(c *fiber.Ctx) error {
	pq := parserBodyQuery(c)
	data := services.CreateQuery(pq)
	return c.JSON(&data)
}
