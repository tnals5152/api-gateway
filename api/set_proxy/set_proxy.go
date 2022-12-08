package setproxy

import (
	"github.com/gofiber/fiber/v2"
	"tnals5152.com/api-gateway/model"
)

func CreateProxy(c *fiber.Ctx) error {

	var path *model.Path

	if err := c.BodyParser(&path); err != nil {
		return c.JSON(err.Error())
	}

	return nil
}
