package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"go-manage/pkg/config"
)

type Config struct {
	rest.RestConf
	DatabaseConf config.DatabaseConf
	RedisConf    redis.RedisConf
	//CasbinConf   casbin.CasbinConf
}
