package config

import (
	"github.com/zeromicro/go-zero/rest"
	"go-manage/pkg/config"
)

type Config struct {
	rest.RestConf
	DatabaseConf config.DatabaseConf
}
