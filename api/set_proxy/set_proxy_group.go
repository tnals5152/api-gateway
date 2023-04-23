package setproxy

import "github.com/gofiber/fiber/v2"

func SetGroup(router fiber.Router) {
	proxy := router.Group("set-proxy")
	{
		// API 추가
		proxy.Post("", CreateProxy)
		// API 수정
	}
}
