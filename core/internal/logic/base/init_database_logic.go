package base

import (
	"context"
	"go-manage/core/internal/svc"
	"go-manage/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InitDatabaseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInitDatabaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitDatabaseLogic {
	return &InitDatabaseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InitDatabaseLogic) InitDatabase() (resp *types.BaseMsgResp, err error) {
	//初始化 table
	err = l.svcCtx.BaseModel.Init(l.ctx)
	if err != nil {
		return nil, err
	}
	resp = &types.BaseMsgResp{
		Code: 200,
		Msg:  "创建成功",
	}
	return
}
