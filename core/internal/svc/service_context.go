package svc

import (
	"github.com/casbin/casbin/v2"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"go-manage/core/internal/config"
	"go-manage/core/internal/middleware"
	"go-manage/pkg/model"
)

type ServiceContext struct {
	Config    config.Config
	Redis     *redis.Redis
	Authority rest.Middleware
	Casbin    *casbin.Enforcer
	UserModel model.UserModel
	BaseModel model.BaseModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	client := model.NewEntClientService(c.DatabaseConf)
	rds := redis.MustNewRedis(c.RedisConf)
	csb := c.CasbinConf.MustNewCasbinWithRedisWatcher(c.DatabaseConf.Type, c.DatabaseConf.GetDSN(), c.RedisConf)
	return &ServiceContext{
		Config:    c,
		Redis:     rds,
		Casbin:    csb,
		BaseModel: model.NewBaseModel(client),
		UserModel: model.NewUserModel(client),
		Authority: middleware.NewAuthorityMiddleware(csb, rds).Handle,
	}
}
