package svc

import (
	"go-manage/core/internal/config"
	"go-manage/pkg/model"
)

type ServiceContext struct {
	Config config.Config
	Db     *model.EntClientSvc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Db:     model.NewEntClientService(c.DatabaseConf),
	}
}
