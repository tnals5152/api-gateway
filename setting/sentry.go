package setting

func SentryInit() {

	// // To initialize Sentry's handler, you need to initialize Sentry itself beforehand
	// if err := sentry.Init(sentry.ClientOptions{
	// 	Dsn:           "https://5110376c3d80838e6625b49c4bb4b60f@o4505828089331712.ingest.sentry.io/4505828090773504",
	// 	EnableTracing: true,
	// 	// Set TracesSampleRate to 1.0 to capture 100%
	// 	// of transactions for performance monitoring.
	// 	// We recommend adjusting this value in production,
	// 	TracesSampleRate: 1.0,
	// }); err != nil {
	// 	fmt.Printf("Sentry initialization failed: %v", err)
	// }

	// // Create an instance of sentryfasthttp
	// sentryHandler := sentryfiber.New(sentryfiber.Options{})

	// // After creating the instance, you can attach the handler as one of your middleware
	// fastHTTPHandler := sentryHandler.Handle(func(ctx *fasthttp.RequestCtx) {
	// 	panic("y tho")
	// })

	// fmt.Println("Listening and serving HTTP on :3000")

	// // And run it
	// if err := fasthttp.ListenAndServe(":3000", fastHTTPHandler); err != nil {
	// 	panic(err)
	// }
}
