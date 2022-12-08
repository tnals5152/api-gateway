package getproxy

import "github.com/gofiber/fiber/v2"

func SetGroup(router fiber.Router) {
	proxy := router.Group("proxy")
	{
		// + 는 최소 1개 path가 전달되어야 한다.
		proxy.All("+", GetProxy)
	}
}
