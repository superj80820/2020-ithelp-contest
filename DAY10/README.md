本文章同時發佈於：

- [Github(包含程式碼)](https://github.com/superj80820/2020-ithelp-contest/blob/master/DAY10)
- [Medium](https://medium.com/%E9%AB%92%E6%A1%B6%E5%AD%90/day10-%E7%A5%9E%E5%A5%87%E7%9A%84-grpc-%E8%AE%93%E4%BD%A0%E6%8A%8A-call-service-%E7%95%B6%E6%88%90%E4%B8%80%E5%80%8B-function-call-%E5%AF%A6%E4%BD%9C%E7%AF%87-a447b3e76f3b)
- [iT 邦幫忙](https://ithelp.ithome.com.tw/articles/10242947)

```
文章為自己的經驗與夥伴整理的內容，設計沒有標準答案，如有可以改進的地方，請告訴我，我會盡我所能的修改，謝謝大家～
```

---

大家好，繼昨天[DAY09](https://github.com/superj80820/2020-ithelp-contest/blob/master/DAY09)的介紹後，gRPC 相信大家已經稍有概念，接下來將介紹四種 gRPC 的傳輸方式:

- 單向傳輸(Unary): Client 單向，Server 單向
- Server 單向串流(Streaming-Server): Client 單向，Server 串流
- Client 單向串流(Streaming-Client): Client 串流，Server 單向
- 雙向串流(Streaming-Bidirectional): Client 串流，Server 串流

會以一個 Golang Client 一個 Golang Server 來講解。

以下講解 code 全部都放在[DAY10-Github](https://github.com/superj80820/2020-ithelp-contest/blob/master/DAY10)

![](https://i.imgur.com/99W9BCD.png)

## 建立 Protobuf schema

如[DAY09](https://github.com/superj80820/2020-ithelp-contest/blob/master/DAY09)所述，必須先建立好一個 Protobuf schema，之後我們就可以用此 schema 產生 Client 與 Server 的 code。

`unary`傳輸

```proto
syntax = "proto3";

package helloworld;

service Greeter {
  rpc SayHello (HelloRequest) returns (HelloReply) {}
}

message HelloRequest {
  string name = 1;
}

message HelloReply {
  string message = 1;
}
```

`streaming-server`傳輸

```proto
syntax = "proto3";

package helloworld;

service Greeter {
  rpc SayHello (HelloRequest) returns (stream HelloReply) {}
}

message HelloRequest {
  string name = 1;
}

message HelloReply {
  string message = 1;
}
```

`streaming-client`傳輸

```proto
syntax = "proto3";

package helloworld;

service Greeter {
  rpc SayHello (stream HelloRequest) returns (HelloReply) {}
}

message HelloRequest {
  string name = 1;
}

message HelloReply {
  string message = 1;
}
```

`streaming-bidirectional`傳輸

```proto
syntax = "proto3";

package helloworld;

service Greeter {
  rpc SayHello (stream HelloRequest) returns (stream HelloReply) {}
}

message HelloRequest {
  string name = 1;
}

message HelloReply {
  string message = 1;
}
```

- `syntax`: 為 protobuf 的版本，目前常見的 proto2 與 proto3，在 require 這個關鍵字上做了很大的變動，詳情改動可以看[設計思考 — Protocol Buffers 3 為什麼](https://medium.com/@leon740727/%E8%A8%AD%E8%A8%88%E6%80%9D%E8%80%83-protocol-buffers-3-%E7%82%BA%E4%BB%80%E9%BA%BC-49219fc87bb7)。
- `package`: 為產出來 code 的 packge name。
- `message`: 為 protobuf 的訊息，你可以想像是 Restful API 的 request 與 response，但是介面更為嚴謹。
- `service`: gRPC 的方法由此定義，以`rpc SayHello (stream HelloRequest) returns (HelloReply) {}`來說，`SayHello`是方法名，`HelloRequest`為參數，`HelloReply`為 result，有無`stream`代表了是否要用串流還是單向。

設計好了之後就可以執行以下 script 產生出`gen`資料夾，裡面就有實際的 code 了。

![](https://i.imgur.com/Fyt9vTE.png)

```bash
$ build-proto.sh
```

## 實際運作

```golang
package main

import (
	"context"
	"log"
	"net"

	pb "unary/gen/grpc-gateway/gen"

	"google.golang.org/grpc"
)

const (
	port = ":50052"
)

// SayHello implements helloworld.GreeterServer.SayHello
func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

// Server ...
type Server struct {
	pb.UnimplementedGreeterServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
```

以 unary 來說，將要`import pb "unary/gen/grpc-gateway/gen"`，並將`pb.UnimplementedGreeterServer`實作，最後透過`grpc.NewServer()`來創建 gRPC Server，並透過`pb.RegisterGreeterServer(s, &Server{})`來註冊實作，最後以`s.Serve(lis)`啟動 Server。

而四種傳輸方法的測試方式都是:

1. `go run server/main.go`: 啟動 server
2. `go run client/main.go`: 使用 client 來呼叫

![](https://i.imgur.com/6Wjs2jw.png)

都大同小異，實作細節可以看 code，我舉出比較要注意的地方，

`srv pb.Greeter_SayHelloServer`為 Streaming 的方式，在 client 端呼叫後:

1. 如果是 Streaming-Client: Server 可以透過此參數以`.Recv()`的方式持續接收資料，Client 可以透過此參數以`.Send()`的方式持續傳送資料。
2. 如果是 Streaming-Server: Server 可以透過此參數以`.Send()`的方式持續傳送資料，Client 可以透過此參數以`.Recv()`的方式持續傳送資料。
3. 如果是 Streaming-Bidirectional: Server 可以透過此參數以`.Send()`的方式持續傳送資料與`.Recv()`的方式持續接收資料，Client 可以透過此參數以`.Send()`的方式持續傳送資料與`.Recv()`的方式持續接收資料。

## 參考

- [Core concepts, architecture and lifecycle – gRPC](https://grpc.io/docs/what-is-grpc/core-concepts/)
- [grpc-go/examples at master · grpc/grpc-go](https://github.com/grpc/grpc-go/tree/master/examples)
- [Basics Tutorial – gRPC](https://grpc.io/docs/languages/go/basics/)
- [go-grpc-examples/client.go at master · itsksaurabh/go-grpc-examples](https://github.com/itsksaurabh/go-grpc-examples/blob/master/stream/bi-directional-streaming/feeds/feedClient/client.go)
- [設計思考 — Protocol Buffers 3 為什麼. 簡單是一件非常難的事! 而深思熟慮過的簡單，可以給我們最多的思考與學習 | by Leon Tsai | Medium](https://medium.com/@leon740727/%E8%A8%AD%E8%A8%88%E6%80%9D%E8%80%83-protocol-buffers-3-%E7%82%BA%E4%BB%80%E9%BA%BC-49219fc87bb7)
- [比起 JSON 更方便、更快速、更簡短的 Protobuf 格式 – 電腦玩瞎咪](https://yami.io/protobuf/)
