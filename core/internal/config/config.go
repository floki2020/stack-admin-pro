package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"go-manage/pkg/config"
	"go-manage/pkg/plugin/casbin"
)

type Config struct {
	rest.RestConf
	DatabaseConf config.DatabaseConf
	RedisConf    redis.RedisConf
	CasbinConf   casbin.CasbinConf
}
