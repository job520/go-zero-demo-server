# go-zero-demo-server

### 示例用法（服务端/客户端）

1. 目录结构：
```
  |—— demo/   // 根目录
  ****|—— api/
  ********|—— client/
  ************|—— main.go
  ********|—— server/
  ************|—— test.api
  ****|—— rpc/
  ********|—— client/
  ************|—— main.go
  ********|—— server/
  ************|—— test.proto
```

2. 文件内容：
  1. `demo/api/server/test.api`
  ```
  type (
    Req {
      Id string `path:"id"`
    }
    Resp {
      Id   string `json:"id"`
      Name string `json:"name"`
    }
  )
  service test {
    @handler testGet
    get /test/:id (Req) returns (Resp)
  }
  ```
  2. `demo/api/client/main.go`
  ```
  package main
  import (
    "fmt"
    "io"
    "io/ioutil"
    "net/http"
  )
  func HTTP_GET(url string) (string, error) {
    resp, err := http.Get(url)
    if resp != nil {
      defer resp.Body.Close()
    }
    if err != nil {
      return "", err
    } else {
      if resp.StatusCode == 200 {
        var bodyReader io.ReadCloser = resp.Body
        body, err := ioutil.ReadAll(bodyReader)
        if err != nil {
          return "", err
        } else {
          return string(body), nil
        }
      } else {
        return "", fmt.Errorf("服务器异常")
      }
    }
  }
  func main() {
    url := "http://localhost:8888/test/1"
    ret, err := HTTP_GET(url)
    fmt.Println(ret, err)
  }
  ```
  3. `demo/rpc/server/test.proto`
  ```
  syntax = "proto3";
  package test;
  option go_package = "./test";
  service Test {
      rpc test(Req) returns(Resp);
  }
  message Req {
      uint32 id = 1;
  }
  message Resp {
      uint32 id = 1;
      string name = 2;
  }
  ```
  4. `demo/rpc/client/main.go`
  ```
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
  ```

3. 初始化go module：
```
cd demo
go mod init demo
```

4. 生成api服务端代码：
```
cd demo/api/server
goctl api go -api test.api -dir .
```

5. 生成rpc服务端代码：
```
cd demo/rpc/server
goctl rpc protoc test.proto --go_out=./types --go-grpc_out=./types --zrpc_out=.
```

6. 代码地址：
[点击访问](https://github.com/job520/go-zero-demo-server)


7. 启动api服务：
```
cd demo/api/server
go run test.go -f etc/test.yaml
```

8. 调用api客户端：
```
cd demo/api/client
go run main.go
```

9. 启动rpc服务：
```
cd demo/rpc/server
go run test.go -f etc/test.yaml
```

10. 调用rpc客户端：
```
cd demo/rpc/client
go run main.go
```
