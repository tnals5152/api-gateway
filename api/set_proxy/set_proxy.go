package setproxy

import (
	"github.com/gofiber/fiber/v2"
	"tnals5152.com/api-gateway/handler"
	"tnals5152.com/api-gateway/model"
)

func CreateProxy(c *fiber.Ctx) error {

	// model.RequestResource 형식으로 들어오도록 변경 예정
	var resource *model.RequestResource

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
