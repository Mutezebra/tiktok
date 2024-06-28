package pack

import (
    "github.com/mutezebra/tiktok/pkg/consts"
    "github.com/mutezebra/tiktok/pkg/inject"
    "github.com/mutezebra/tiktok/pkg/oss"
    "github.com/mutezebra/tiktok/pkg/utils"
    "github.com/mutezebra/tiktok/user/config"
    user "github.com/mutezebra/tiktok/user/domain/service"
    "github.com/mutezebra/tiktok/user/interface/persistence/database"
    "github.com/mutezebra/tiktok/user/usecase"
)

func UserRegistry() *inject.Registry {
    return &inject.Registry{
        Addr:   config.Conf.Service[consts.UserServiceName].Address,
        Key:    config.Conf.Service[consts.UserServiceName].ServiceName + "/" + config.Conf.Service[consts.UserServiceName].Address,
        Prefix: config.Conf.Etcd.ServicePrefix}
}

func ApplyUser() *usecase.UserCase {
    repo := database.NewUserRepository()
    ossModel := oss.NewOssModel()
    mfaModel := utils.NewMFAModel()

    service := &user.Service{
        Repo: repo,
        MFA:  mfaModel,
        OSS:  ossModel,
    }

    return usecase.NewUserUseCase(repo, user.NewService(service))
}
