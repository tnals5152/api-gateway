package getproxy

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	_ "github.com/swaggo/fiber-swagger/example/docs"
	constant "tnals5152.com/api-gateway/const"
	"tnals5152.com/api-gateway/handler"
)

func GetProxyByFunctionName(c *fiber.Ctx) error {
	// 1. functionName 추출
	functionName := c.Params("functionName")
	// 2. 일치하는 resource를 DB에서 조회
	respnse, err := new(handler.ContextHandler).
		SetCtx(c).
		CallByFunctionName().
		GetCorrectResource(functionName).
		Call()

	if err != nil {
		// TODO: response Model 정의하기
		return c.JSON(map[string]any{
			"error": err.Error(),
		})
	}

	return c.JSON(respnse)
}

// @Summary Host information collection
func GetProxy(c *fiber.Ctx) error {
	// 1. parameters를 자른다.
	params := c.Params(constant.PLUS)
	pathParams := strings.Split(params, constant.SLASH)
	// 2. request api의 path로 DB에서 일치하는 endpoint_path를 찾는다.
	respnse, err := new(handler.ContextHandler).
		SetCtx(c).
		SetReqeustParams(pathParams).
		GetCorrectResource().
		Call()

	if err != nil {
		// TODO: response Model 정의하기
		return c.JSON(map[string]any{
			"error": err.Error(),
		})
	}

	return c.JSON(respnse)
}
