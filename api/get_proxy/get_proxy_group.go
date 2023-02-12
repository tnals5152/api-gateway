package getproxy

import "github.com/gofiber/fiber/v2"

func SetGroup(router fiber.Router) {
	proxy := router.Group("get-proxy")
	{
		// + 는 최소 1개 path가 전달되어야 한다.
		// request_path로 요청되는 api만 해당된다.
		proxy.All("+", GetProxy)

		// path collection의 name으로 호출될 때 사용한다.
	}
}
