package pack

import (
	interaction "github.com/mutezebra/tiktok/interaction/domain/service"
	"github.com/mutezebra/tiktok/interaction/interface/persistence/database"

	"github.com/mutezebra/tiktok/interaction/config"
	"github.com/mutezebra/tiktok/interaction/usecase"
	"github.com/mutezebra/tiktok/pkg/consts"
	"github.com/mutezebra/tiktok/pkg/inject"
)

func InteractionRegistry() *inject.Registry {
	return &inject.Registry{
		Addr:   config.Conf.Service[consts.InteractionServiceName].Address,
		Key:    config.Conf.Service[consts.InteractionServiceName].ServiceName + "/" + config.Conf.Service[consts.InteractionServiceName].Address,
		Prefix: config.Conf.Etcd.ServicePrefix}
}

func ApplyInteraction() *usecase.InteractionCase {
	repo := database.NewInteractionRepository()
	service := &interaction.Service{
		Repo: repo,
	}
	return usecase.NewInteractionCase(repo, interaction.NewService(service))
}
