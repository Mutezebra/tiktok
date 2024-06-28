package pack

import (
	relation "github.com/mutezebra/tiktok/relation/domain/service"
	"github.com/mutezebra/tiktok/relation/interface/persistence/database"
	"github.com/mutezebra/tiktok/relation/usecase"
	"github.com/mutezebra/tiktok/relation/config"
	"github.com/mutezebra/tiktok/pkg/consts"
	"github.com/mutezebra/tiktok/pkg/inject"
)

func RelationRegistry() *inject.Registry {
	return &inject.Registry{
		Addr:   config.Conf.Service[consts.RelationServiceName].Address,
		Key:    config.Conf.Service[consts.RelationServiceName].ServiceName + "/" + config.Conf.Service[consts.RelationServiceName].Address,
		Prefix: config.Conf.Etcd.ServicePrefix}
}

func ApplyRelation() *usecase.RelationCase {
	repo := database.NewRelationRepository()
	service := relation.NewService(repo)
	return usecase.NewRelationCase(service, repo)
}
