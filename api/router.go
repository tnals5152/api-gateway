package api

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	"github.com/spf13/viper"
	getproxy "tnals5152.com/api-gateway/api/get_proxy"
	setproxy "tnals5152.com/api-gateway/api/set_proxy"
	constant "tnals5152.com/api-gateway/const"
	ct "tnals5152.com/api-gateway/const"
	"tnals5152.com/api-gateway/utils"

	sentryfiber "github.com/aldy505/sentry-fiber"
	"github.com/getsentry/sentry-go"
)

func SetRouter() {
	app := fiber.New(fiber.Config{
		ReadTimeout:  utils.GetTimeout(ct.ServerReadTimeout) * time.Second,
		WriteTimeout: utils.GetTimeout(ct.ServerWriteTimeout) * time.Second,
		AppName:      viper.GetString(ct.ServerAppName),
	})

	app.Use(cors.New())
	app.Use(SentryInit())
	app.Use(logger.New(logger.Config{
		Output:     utils.LogFile,
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Asia/Seoul",
	}))

	router := app.Group("api")

	setproxy.SetGroup(router)
	getproxy.SetGroup(router)

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "http://example.com/doc.json",
		DeepLinking: false,
		// Expand ("list") or Collapse ("none") tag groups by default
		DocExpansion: "none",
		// Prefill OAuth ClientId on Authorize popup
		OAuth: &swagger.OAuthConfig{
			AppName:  "OAuth Provider",
			ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
		},
		// Ability to change OAuth2 redirect uri location
		OAuth2RedirectUrl: "http://localhost:8888/swagger/oauth2-redirect.html",
	}))

	app.Listen(viper.GetString(ct.ServerPort))
}

func SentryInit() func(*fiber.Ctx) error {

	// To initialize Sentry's handler, you need to initialize Sentry itself beforehand
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:           viper.GetString(constant.SENTRY_DSN),
		EnableTracing: true,
		// Set TracesSampleRate to 1.0 to capture 100%
		// of transactions for performance monitoring.
		// We recommend adjusting this value in production,
		TracesSampleRate: 1.0,
		Debug:            true,
		AttachStacktrace: true,
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v", err)
	}

	// Create an instance of sentryfasthttp
	sentryHandler := sentryfiber.New(sentryfiber.Options{})

	return sentryHandler
}
