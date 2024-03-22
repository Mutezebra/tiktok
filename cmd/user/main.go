package main

import (
	"context"
	"github.com/Mutezebra/tiktok/cmd/user/pack"
	"net"

	"github.com/Mutezebra/tiktok/app/interface/persistence/database"

	"github.com/cloudwego/kitex/server"

	"github.com/Mutezebra/tiktok/config"
	"github.com/Mutezebra/tiktok/consts"
	"github.com/Mutezebra/tiktok/kitex_gen/api/user/userservice"
	"github.com/Mutezebra/tiktok/pkg/inject"
	"github.com/Mutezebra/tiktok/pkg/log"
)

func main() {
	UserInit()
	registry, err := inject.NewRegistry(pack.UserRegistry())
	if err != nil {
		log.LogrusObj.Panic(err)
	}

	err = registry.Register(context.Background())
	defer registry.Close()
	if err != nil {
		log.LogrusObj.Panic(err)
	}

	addr, _ := net.ResolveTCPAddr("tcp", config.Conf.Service[consts.UserServiceName].Address)
	srv := userservice.NewServer(inject.ApplyUser(), server.WithServiceAddr(addr))
	err = srv.Run()
	if err != nil {
		log.LogrusObj.Panic(err)
	}
}

func UserInit() {
	config.InitConfig()
	log.InitLog()
	database.InitMysql()
}
