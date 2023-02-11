package getproxy

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	constant "tnals5152.com/api-gateway/const"
	"tnals5152.com/api-gateway/handler"
)

func GetProxy(c *fiber.Ctx) error {
	// 1. parameters를 자른다.
	params := c.Params(constant.PLUS)
	pathParams := strings.Split(params, constant.SLASH)
	// 2. request api의 path로 DB에서 일치하는 endpoint_path를 찾는다.
	handler.GetEndpointPath(pathParams)

	return nil
}
