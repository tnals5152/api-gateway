package setproxy

import (
	"github.com/gofiber/fiber/v2"
	host_pkg "tnals5152.com/api-gateway/db/query"
	"tnals5152.com/api-gateway/handler"
	"tnals5152.com/api-gateway/model"
)

func CreateProxy(c *fiber.Ctx) error {

	var resource *model.Resource

	if err := c.BodyParser(&resource); err != nil {
		return c.JSON(err.Error())
	}

	// 필수 필드 검증 추가
	if err := resource.Validate(); err != nil {
		return c.JSON(err.Error())
	}

	err := handler.SetProxyData(resource)

	if err != nil {
		return c.JSON(err.Error())
	}

	return nil
}
