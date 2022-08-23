// Code generated by goctl. DO NOT EDIT!
// Source: test.proto

package server

import (
	"context"

	"demo/rpc/server/internal/logic"
	"demo/rpc/server/internal/svc"
	"demo/rpc/server/types/test"
)

type TestServer struct {
	svcCtx *svc.ServiceContext
	test.UnimplementedTestServer
}

func NewTestServer(svcCtx *svc.ServiceContext) *TestServer {
	return &TestServer{
		svcCtx: svcCtx,
	}
}

func (s *TestServer) Test(ctx context.Context, in *test.Req) (*test.Resp, error) {
	l := logic.NewTestLogic(ctx, s.svcCtx)
	return l.Test(in)
}