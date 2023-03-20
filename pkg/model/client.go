package model

import (
	"github.com/zeromicro/go-zero/core/logx"
	"go-manage/pkg/config"
	"go-manage/pkg/ent"
)

func NewEntClientService(c config.DatabaseConf) *ent.Client {
	db := ent.NewClient(
		ent.Debug(),
		ent.Log(logx.Info),
		ent.Driver(c.NewNoCacheDriver()),
	)
	return db
}
