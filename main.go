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

func main() {
	defer utils.LogFile.Close()

	api.SetRouter()
}
