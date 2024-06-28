package main

import (
    "context"
    "net"

    "github.com/mutezebra/tiktok/video/cmd/pack"
    "github.com/mutezebra/tiktok/video/interface/persistence/cache"
    "github.com/mutezebra/tiktok/video/interface/persistence/database"

    "github.com/cloudwego/kitex/pkg/transmeta"

    "github.com/mutezebra/tiktok/pkg/kitex_gen/api/video/videoservice"

    "github.com/mutezebra/tiktok/pkg/oss"

    "github.com/cloudwego/kitex/server"

    "github.com/mutezebra/tiktok/pkg/consts"
    "github.com/mutezebra/tiktok/pkg/inject"
    "github.com/mutezebra/tiktok/pkg/log"
    "github.com/mutezebra/tiktok/video/config"
)

func main() {
    VideoInit()
    registry := inject.NewRegistry(pack.VideoRegistry())
    defer registry.Close()
    registry.MustRegister(context.Background())

    addr, _ := net.ResolveTCPAddr("tcp", config.Conf.Service[consts.VideoServiceName].Address)
    srv := videoservice.NewServer(
        pack.ApplyVideo(),
        server.WithServiceAddr(addr),
        server.WithMetaHandler(transmeta.ServerTTHeaderHandler))
    err := srv.Run()
    if err != nil {
        log.LogrusObj.Panic(err)
    }
}

func VideoInit() {
    config.InitConfig()
    log.InitLog()
    database.InitMysql()
    cache.InitCache()
    oss.InitOSS()
}
