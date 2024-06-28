package main

import (
	"context"
	"fmt"
	"net"

	"github.com/mutezebra/tiktok/user/cmd/pack"
	"github.com/mutezebra/tiktok/user/interface/persistence/database"

	"github.com/cloudwego/kitex/pkg/transmeta"

	"github.com/mutezebra/tiktok/pkg/oss"

	"github.com/cloudwego/kitex/server"

	"github.com/mutezebra/tiktok/pkg/consts"
	"github.com/mutezebra/tiktok/pkg/inject"
	"github.com/mutezebra/tiktok/pkg/kitex_gen/api/user/userservice"
	"github.com/mutezebra/tiktok/pkg/log"
	"github.com/mutezebra/tiktok/user/config"
)

func main() {
	UserInit()
	fmt.Println("have init")
	registry := inject.NewRegistry(pack.UserRegistry())
	defer registry.Close()
	registry.MustRegister(context.Background())

	addr, _ := net.ResolveTCPAddr("tcp", config.Conf.Service[consts.UserServiceName].Address)
	srv := userservice.NewServer(
		pack.ApplyUser(),
		server.WithServiceAddr(addr),
		server.WithMetaHandler(transmeta.ServerTTHeaderHandler))
	err := srv.Run()
	if err != nil {
		log.LogrusObj.Panic(err)
	}
}

func UserInit() {
	config.InitConfig()
	log.InitLog()
	database.InitMysql()
	oss.InitOSS()
}
