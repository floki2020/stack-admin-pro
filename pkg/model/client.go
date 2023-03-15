package model

import (
	"github.com/zeromicro/go-zero/core/logx"
	"go-manage/pkg/config"
	"go-manage/pkg/ent"
)

// EntClientSvc TODO: Returns the object directly
type EntClientSvc struct {
	Db *ent.Client
}

func NewEntClientService(c config.DatabaseConf) *EntClientSvc {
	db := ent.NewClient(
		ent.Debug(),
		ent.Log(logx.Info),
		ent.Driver(c.NewNoCacheDriver()),
	)
	return &EntClientSvc{
		Db: db,
	}
}
