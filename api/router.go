package api

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/spf13/viper"
	getproxy "tnals5152.com/api-gateway/api/get_proxy"
	setproxy "tnals5152.com/api-gateway/api/set_proxy"
	ct "tnals5152.com/api-gateway/const"
	"tnals5152.com/api-gateway/utils"
)

func SetRouter() {
	app := fiber.New(fiber.Config{
		ReadTimeout:  utils.GetTimeout(ct.ServerReadTimeout) * time.Second,
		WriteTimeout: utils.GetTimeout(ct.ServerWriteTimeout) * time.Second,
		AppName:      viper.GetString(ct.ServerAppName),
	})

	app.Use(cors.New())
	app.Use(logger.New(logger.Config{
		Output:     utils.LogFile,
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Asia/Seoul",
	}))

	router := app.Group("api")

	setproxy.SetGroup(router)
	getproxy.SetGroup(router)

	app.Listen(viper.GetString(ct.ServerPort))
}
