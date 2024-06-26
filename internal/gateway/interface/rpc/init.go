package rpc

import (
	"context"
	"fmt"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/streamclient"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"

	"github.com/mutezebra/tiktok/gateway/config"
	"github.com/mutezebra/tiktok/pkg/consts"
	"github.com/mutezebra/tiktok/pkg/inject"
	"github.com/mutezebra/tiktok/pkg/kitex_gen/api/interaction/interactionservice"
	"github.com/mutezebra/tiktok/pkg/kitex_gen/api/relation/relationservice"
	"github.com/mutezebra/tiktok/pkg/kitex_gen/api/user/userservice"
	"github.com/mutezebra/tiktok/pkg/kitex_gen/api/video/videoservice"
	"github.com/mutezebra/tiktok/pkg/log"
)

var (
	err  error
	Conf *config.Config

	Resolver          *inject.Resolver
	UserClient        userservice.Client
	VideoClient       videoservice.Client
	VideoStreamClient videoservice.StreamClient
	InteractionClient interactionservice.Client
	RelationClient    relationservice.Client
)

func Init() {
	Conf = config.Conf
	Resolver, err = inject.NewResolver()
	if err != nil {
		log.LogrusObj.Panic(err)
	}

	initClient(consts.UserServiceName)
	// initClient(consts.VideoServiceName)
	// initClient(consts.InteractionServiceName)
	// initClient(consts.RelationServiceName)
}

func initClient(serviceName string) {
	switch serviceName {
	case consts.UserServiceName:
		UserClient = userservice.MustNewClient(serviceName,
			client.WithHostPorts(discovery(serviceName)...),
			client.WithTransportProtocol(transport.TTHeader),
			client.WithMetaHandler(transmeta.ClientTTHeaderHandler))
	case consts.VideoServiceName:
		VideoClient = videoservice.MustNewClient(serviceName,
			client.WithHostPorts(discovery(serviceName)...),
			client.WithTransportProtocol(transport.TTHeader),
			client.WithMetaHandler(transmeta.ClientTTHeaderHandler))
		VideoStreamClient = videoservice.MustNewStreamClient(serviceName,
			streamclient.WithHostPorts(discovery(serviceName)...),
			streamclient.WithMetaHandler(transmeta.ClientTTHeaderHandler))
	case consts.InteractionServiceName:
		InteractionClient = interactionservice.MustNewClient(serviceName,
			client.WithHostPorts(discovery(serviceName)...),
			client.WithTransportProtocol(transport.TTHeader),
			client.WithMetaHandler(transmeta.ClientTTHeaderHandler))
	case consts.RelationServiceName:
		RelationClient = relationservice.MustNewClient(serviceName,
			client.WithHostPorts(discovery(serviceName)...),
			client.WithTransportProtocol(transport.TTHeader),
			client.WithMetaHandler(transmeta.ClientTTHeaderHandler),
		)
	}
}

func discovery(serviceName string) []string {
	prefix := Conf.Etcd.ServicePrefix + Conf.Service[serviceName].ServiceName
	addr, err := Resolver.ResolveWithPrefix(context.Background(), prefix)
	fmt.Println(addr)
	if err != nil {
		log.LogrusObj.Panic(err)
	}
	return addr
}
