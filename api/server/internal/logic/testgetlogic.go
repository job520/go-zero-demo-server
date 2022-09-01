package logic

import (
	"context"
	"gorm.io/gorm"

	"demo/api/server/internal/svc"
	"demo/api/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TestGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTestGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TestGetLogic {
	return &TestGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func (l *TestGetLogic) TestGet(req *types.Req) (resp *types.Resp, err error) {
	db := l.svcCtx.Db
	rdb := l.svcCtx.Redis
	var ctx = context.Background()
	// 使用 gorm
	db.AutoMigrate(&Product{})
	// 使用 redis
	err = rdb.Set(ctx, "foo", "bar", 0).Err()
	if err != nil {
		logx.Error(err)
	}
	val, err := rdb.Get(ctx, "foo").Result()
	if err != nil {
		logx.Error(err)
	}
	logx.Info("foo: ", val)
	return &types.Resp{
		Id:   "1",
		Name: l.svcCtx.Config.TestValue,
	}, nil
}
