package logic

import (
	"context"
	"gorm.io/gorm"

	"demo/rpc/server/internal/svc"
	"demo/rpc/server/types/test"

	"github.com/zeromicro/go-zero/core/logx"
)

type TestLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TestLogic {
	return &TestLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func (l *TestLogic) Test(in *test.Req) (*test.Resp, error) {
	db := l.svcCtx.Db
	rdb := l.svcCtx.Redis
	var ctx = context.Background()
	// 使用 gorm
	db.AutoMigrate(&Product{})
	// 使用 redis
	err := rdb.Set(ctx, "foo", "bar", 0).Err()
	if err != nil {
		logx.Error(err)
	}
	val, err := rdb.Get(ctx, "foo").Result()
	if err != nil {
		logx.Error(err)
	}
	logx.Info("foo: ", val)
	return &test.Resp{}, nil
}
