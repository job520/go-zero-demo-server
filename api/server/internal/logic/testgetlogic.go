package logic

import (
	"context"

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

func (l *TestGetLogic) TestGet(req *types.Req) (resp *types.Resp, err error) {
	// todo: add your logic here and delete this line

	return &types.Resp{
		Id:   "",
		Name: "",
	}, nil
}
