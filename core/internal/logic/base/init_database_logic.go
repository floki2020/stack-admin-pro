package base

import (
	"context"
	"entgo.io/ent/dialect/sql/schema"
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
	l.svcCtx.Db.Db.Schema.Create(l.ctx, schema.WithForeignKeys(false), schema.WithDropColumn(true),
		schema.WithDropIndex(true))
	_, err = l.svcCtx.Db.Db.User.Create().SetUserName("admin").SetPassword("123").Save(l.ctx)

	resp = &types.BaseMsgResp{
		Code: 200,
		Msg:  "创建成功",
	}
	return
}
