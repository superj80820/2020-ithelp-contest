本文章同時發佈於：

- [Github(包含程式碼)](https://github.com/superj80820/2020-ithelp-contest/blob/master/DAY13)
- [Medium](https://medium.com/%E9%AB%92%E6%A1%B6%E5%AD%90/day13-clean-architecture-%E7%9A%84%E5%8A%9B%E9%87%8F-%E7%84%A1%E7%97%9B%E5%BE%9E-restful-api-%E8%BD%89%E6%8F%9B%E6%88%90-grpc-server-e1b64346aa14)
- [iT 邦幫忙](https://ithelp.ithome.com.tw/articles/10245345)

```
文章為自己的經驗與夥伴整理的內容，設計沒有標準答案，如有可以改進的地方，請告訴我，我會盡我所能的修改，謝謝大家～
```

---

大家好，還記得在 DAY6~DAY8 的 Clean Architecture 篇，使用了 Restful API 來設計，但是後來的介紹卻都是 gRPC，其實最大的目的是要介紹本篇，將 Restful API Server 換成 gRPC Server，而使用 Clean Architecture 你可以

> 讓換框架或引擎只需更動 delivery 層，不引響商務邏輯(Business Logic)

如果對原本的 Server 設計不清楚，可以先閱讀以下文章:

- [DAY6 - 你的 Backend 可以更有彈性一點 - Clean Architecture 概念篇](https://github.com/superj80820/2020-ithelp-contest/blob/master/DAY06)
- [DAY7 - 奔放的 Golang，卻隱藏著有紀律的架構！ - Clean Architecture 實作篇](https://github.com/superj80820/2020-ithelp-contest/blob/master/DAY07)
- [DAY8 - 讓你的 Backend 萬物皆虛，萬事皆可測 - Clean Architecture 測試篇](https://github.com/superj80820/2020-ithelp-contest/blob/master/DAY08)

## 怎麼 run 起來？

建議大家直接先把 Server run 起來，這樣會比較有感覺，

請先 clone [Github-Example-Code](https://github.com/superj80820/2020-ithelp-contest)

```bash
$ cd DAY13/go-server
$ docker-compose up
```

![](https://i.imgur.com/47oRFAw.png)

並且安裝[Bloomrpc](https://github.com/uw-labs/bloomrpc/releases)來測試，他類似於 gRPC 界的 Postman。(當然你也可以寫 gRPC client 來測試，不過這邊以有介面的 tool 來測試)

安裝好後，可以在左上方的`+按鈕`直接 import protobuf 的 schema，

![](https://i.imgur.com/IRVx7kr.png)

修改好 gRPC Server 的 URL 成`localhost:6000`後，點選`綠色test`按鈕，即可看到 gRPC Server 已經處理完你的請求，並且回傳 response。

![](https://i.imgur.com/JGhVBLw.png)

Ｗ ow，可以使用 gRPC Server 了，接下來將介紹實作，

![](https://i.imgur.com/QDH05Tt.png)

真的有那麼少？！看看就知道 XD。

## 實作 gRPC schema

```proto
syntax = "proto3";

package digimon;

service Digimon {
  rpc Create (CreateRequest) returns (CreateResponse) {}
  rpc Query (QueryRequest) returns (QueryResponse) {}
  rpc Foster (FosterRequest) returns (FosterResponse) {}
}

// Requests

message CreateRequest {
  string name = 1;
}

message QueryRequest {
  string id = 1;
}

message FosterRequest {
  message Food {
    string name = 1;
  }

  string id = 1;
  Food food = 2;
}

// Responses

message CreateResponse {
  string id = 1;
  string name = 2;
  string status = 3;
}

message QueryResponse {
  string id = 1;
  string name = 2;
  string status = 3;
}

message FosterResponse {
}
```

在 Digimon service 裡分別新增 3 個 method，分別是`Create`、`Query`、`Foster`，你可以把他理解為 Restful API 的三個 API，並且分別給 3 個 method 定義好 request 與 response，在 code 中必須符合此 protobuf schema 來實作。

## 實作 gRPC Server

在 digimon/delivery 資料夾新增一個 grpc 資料夾，他代表了另一種引擎的 deliver，

![](https://i.imgur.com/CQ0hVDM.png)

實作`grpc_handler.go`，

```golang
// go-server/digimon/delivery/grpc/grpc_handler.go

// ... 其他程式碼

// DigimonHandler ...
type DigimonHandler struct {
	DigimonUsecase domain.DigimonUsecase
	DietUsecase    domain.DietUsecase
	pb.UnimplementedDigimonServer
}

// NewDigimonHandler ...
func NewDigimonHandler(s *grpcLib.Server, digimonUsecase domain.DigimonUsecase, dietUsecase domain.DietUsecase) {
	handler := &DigimonHandler{
		DigimonUsecase: digimonUsecase,
		DietUsecase:    dietUsecase,
	}

	pb.RegisterDigimonServer(s, handler)
}

// Create ...
func (d *DigimonHandler) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	aDigimon := domain.Digimon{
		Name: req.GetName(),
	}
	if err := d.DigimonUsecase.Store(ctx, &aDigimon); err != nil {
		logrus.Error(err)
		return nil, status.Errorf(codes.Internal, "Internal error. Store failed")
	}

	return &pb.CreateResponse{
		Id:     aDigimon.ID,
		Name:   aDigimon.Name,
		Status: aDigimon.Status,
	}, nil
}

// ... 其他程式碼
```

在`Create()`中，使用`req.GetName()`獲得 Name，並以`DigimonUsecase`的`Store()`來儲存，與原本 Restful API 的 code 做比較，

```golang
// go-server/digimon/delivery/http/digimon_handler.go

// ... 其他程式碼

// DigimonHandler ...
type DigimonHandler struct {
	DigimonUsecase domain.DigimonUsecase
	DietUsecase    domain.DietUsecase
}

// NewDigimonHandler ...
func NewDigimonHandler(e *gin.Engine, digimonUsecase domain.DigimonUsecase, dietUsecase domain.DietUsecase) {
	handler := &DigimonHandler{
		DigimonUsecase: digimonUsecase,
		DietUsecase:    dietUsecase,
	}

	e.GET("/api/v1/digimons/:digimonID", handler.GetDigimonByDigimonID)
	e.POST("/api/v1/digimons", handler.PostToCreateDigimon)
	e.POST("/api/v1/digimons/:digimonID/foster", handler.PostToFosterDigimon)
}

// PostToCreateDigimon ...
func (d *DigimonHandler) PostToCreateDigimon(c *gin.Context) {
	var body swagger.DigimonInfoRequest
	if err := c.BindJSON(&body); err != nil {
		logrus.Error(err)
		c.JSON(500, &swagger.ModelError{
			Code:    3000,
			Message: "Internal error. Parsing failed",
		})
		return
	}
	aDigimon := domain.Digimon{
		Name: body.Name,
	}
	if err := d.DigimonUsecase.Store(c, &aDigimon); err != nil {
		logrus.Error(err)
		c.JSON(500, &swagger.ModelError{
			Code:    3000,
			Message: "Internal error. Store failed",
		})
		return
	}

	c.JSON(200, swagger.DigimonInfo{
		Id:     aDigimon.ID,
		Name:   aDigimon.Name,
		Status: aDigimon.Status,
	})
}

// ... 其他程式碼
```

> 除了交付參數給 Usecase 的方式不同，其他邏輯完全相同！

對的沒錯，在 Restful API 中需要以`swagger.DigimonInfoRequest`搭配`c.BindJSON()`來獲得 JSON Body，最後也以`DigimonUsecase`的`Store()`來儲存，所以，

> Delivery 層單純就是在實作如何把引擎的資訊交付給 Usecase 層

---

很實用對吧！雖然當我知道這樣的架構的時候，我內心偏向下圖 XD:

![](https://i.imgur.com/6BnCS83.png)

(~~滿滿的重構時間~~)

謝謝你的閱讀～

## 參考

- [bloomrpc](https://github.com/uw-labs/bloomrpc)
