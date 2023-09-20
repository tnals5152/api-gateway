package main

import (
	"runtime"

	"tnals5152.com/api-gateway/api"
	"tnals5152.com/api-gateway/db"
	"tnals5152.com/api-gateway/setting"
	"tnals5152.com/api-gateway/utils"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	setting.SetConfig()
	utils.SetLog()
	db.ConnectDB()

}

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8888
// @BasePath /
func main() {
	defer utils.LogFile.Close()
	go setting.SentryInit()

	api.SetRouter()
}
