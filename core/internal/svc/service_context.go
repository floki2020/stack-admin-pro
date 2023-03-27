package svc

import (
	"github.com/casbin/casbin/v2"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"go-manage/core/internal/config"
	"go-manage/pkg/model"
)

type ServiceContext struct {
	Config    config.Config
	Redis     *redis.Redis
	Casbin    *casbin.Enforcer
	UserModel model.UserModel
	BaseModel model.BaseModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	client := model.NewEntClientService(c.DatabaseConf)
	redis := redis.MustNewRedis(c.RedisConf)
	casbin := c.CasbinConf.MustNewCasbinWithRedisWatcher(c.DatabaseConf.Type, c.DatabaseConf.GetDSN(), c.RedisConf)
	return &ServiceContext{
		Config:    c,
		Redis:     redis,
		Casbin:    casbin,
		BaseModel: model.NewBaseModel(client),
		UserModel: model.NewUserModel(client),
	}
}
