package main

import (
	"github.com/mutezebra/tiktok/gateway/config"
	"github.com/mutezebra/tiktok/gateway/interface/persistence/database"
	"github.com/mutezebra/tiktok/gateway/interface/router"
	"github.com/mutezebra/tiktok/pkg/log"
)

func main() {
	GatewayInit()
	h := router.NewRouter()
	h.Spin()
}

func GatewayInit() {
	config.InitConfig()
	log.InitLog(config.Conf.System.Status, config.Conf.System.OS)
	database.InitMysql()
	// rpc.Init()
}
