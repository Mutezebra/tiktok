package pack

import (
	"github.com/mutezebra/tiktok/pkg/consts"
	"github.com/mutezebra/tiktok/pkg/inject"
	"github.com/mutezebra/tiktok/pkg/oss"
	"github.com/mutezebra/tiktok/video/config"
	video "github.com/mutezebra/tiktok/video/domain/service"
	"github.com/mutezebra/tiktok/video/interface/persistence/cache"
	"github.com/mutezebra/tiktok/video/interface/persistence/database"
	"github.com/mutezebra/tiktok/video/usecase"
)

func VideoRegistry() *inject.Registry {
	return &inject.Registry{
		Addr:   config.Conf.Service[consts.VideoServiceName].Address,
		Key:    config.Conf.Service[consts.VideoServiceName].ServiceName + "/" + config.Conf.Service[consts.VideoServiceName].Address,
		Prefix: config.Conf.Etcd.ServicePrefix}
}

func ApplyVideo() *usecase.VideoCase {
	repo := database.NewVideoRepository()
	ossModel := oss.NewOssModel()
	cacheRepo := cache.NewVideoCacheRepo()
	service := &video.Service{
		EnablePopularVideoRank:  true,
		EnableTimedRefreshViews: true,
		Repo:                    repo,
		Cache:                   cacheRepo,
		Oss:                     ossModel,
	}
	return usecase.NewVideoUseCase(repo, cacheRepo, video.NewService(service))
}
