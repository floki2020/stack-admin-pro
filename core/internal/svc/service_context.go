package svc

import (
	"go-manage/core/internal/config"
	"go-manage/pkg/model"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
	BaseModel model.BaseModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	client := model.NewEntClientService(c.DatabaseConf)
	return &ServiceContext{
		Config:    c,
		BaseModel: model.NewBaseModel(client),
		UserModel: model.NewUserModel(client),
	}
}
