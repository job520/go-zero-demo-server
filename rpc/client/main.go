package main

import (
	"context"
	"demo/rpc/server/testclient"
	"fmt"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/zrpc"
)

func main() {
	client := zrpc.MustNewClient(zrpc.RpcClientConf{
		Etcd: discov.EtcdConf{
			Hosts: []string{"127.0.0.1:2379"},
			Key:   "test.rpc",
		},
	})
	t := testclient.NewTest(client)
	resq, err := t.Test(context.Background(), &testclient.Req{Id: 1})
	fmt.Println(resq, err)
}
