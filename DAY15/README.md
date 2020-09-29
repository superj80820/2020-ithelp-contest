本文章同時發佈於：

- [Github(包含程式碼)](https://github.com/superj80820/2020-ithelp-contest/blob/master/DAY15)
- [Medium]()
- [iT 邦幫忙]()

```
文章為自己的經驗與夥伴整理的內容，設計沒有標準答案，如有可以改進的地方，請告訴我，我會盡我所能的修改，謝謝大家～
```

---

大家好，實作完微服務的溝通後，接下來要來實作前端與後端的 gRPC-Web 部分，整體的服務終於有點微服務的樣子了 XD

整體架構如下，除了外部的`Web`前端以外，還需要`Envoy`來轉換 HTTP1.1 至 HTTP2(詳細原因在[DAY11](https://github.com/superj80820/2020-ithelp-contest/tree/master/DAY11))

[//]: #"./digimon-service.drawio.png"

![](https://i.imgur.com/LpNiBo4.png)

## 實際運作

gRPC schemas:

- [gRPC.Digimon](https://github.com/superj80820/2020-ithelp-contest/blob/master/DAY15/schemas/digimon/schema.proto)
- [gRPC.Weather](https://github.com/superj80820/2020-ithelp-contest/blob/master/DAY15/schemas/weather/schema.proto)

以下流程可以與上方 schema 做對照，會比較清楚，

[![](https://mermaid.ink/img/eyJjb2RlIjoiXG5zZXF1ZW5jZURpYWdyYW1cblx0XHRwYXJ0aWNpcGFudCBjIGFzIFdlYlxuXHRcdHBhcnRpY2lwYW50IGUgYXMgRW52b3lcbiAgICBwYXJ0aWNpcGFudCBkIGFzIERpZ2ltb24tU2VydmljZVxuICAgIHBhcnRpY2lwYW50IHcgYXMgV2VhdGhlci1TZXJ2aWNlXG4gICAgYy0-PmU6IGdSUEMuRGlnaW1vbi5DcmVhdGU8YnI-5Ym15bu65pW456K85424XG5cdFx0ZS0-PmQ6IOi9ieeZvFxuXHRcdGQtLT4-ZTog6L2J55m8XG4gICAgZS0tPj5jOiDlm57lgrPmlbjnorznjbjos4foqIpcbiAgICBsb29wIFNlcnZlci1TdHJlYW3kuLLmtYFcbiAgICAgIGMtPj5lOiBnUlBDLkRpZ2ltb24uUXVlcnlTdHJlYW08YnI-5Lul5pW456K85424SUTmkojlj5bmlbjnorznjbjos4foqIpcblx0XHRcdGUtPj5kOiDovYnnmbxcbiAgICAgIGxvb3Ag6ZuZ5ZCR5Liy5rWBXG4gICAgICAgIGQtPj53OiBnUlBDLldlYXRoZXIuUXVlcnk8YnI-5pKI5Y-W5q2k5pW456K85424J-S9jee9rkEn55qE5aSp5rCjXG4gICAgICAgIHctLT4-ZDog5Zue5YKz5aSp5rCjXG4gICAgICBlbmRcbiAgICAgIGQtLT4-ZTog5Zue5YKz5pW456K8542455qE6LOH6KiK6IiH5aSp5rCjXG5cdFx0XHRlLS0-PmM6IOi9ieeZvFxuICAgIGVuZFxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifX0)](https://mermaid-js.github.io/mermaid-live-editor/#/edit/eyJjb2RlIjoiXG5zZXF1ZW5jZURpYWdyYW1cblx0XHRwYXJ0aWNpcGFudCBjIGFzIFdlYlxuXHRcdHBhcnRpY2lwYW50IGUgYXMgRW52b3lcbiAgICBwYXJ0aWNpcGFudCBkIGFzIERpZ2ltb24tU2VydmljZVxuICAgIHBhcnRpY2lwYW50IHcgYXMgV2VhdGhlci1TZXJ2aWNlXG4gICAgYy0-PmU6IGdSUEMuRGlnaW1vbi5DcmVhdGU8YnI-5Ym15bu65pW456K85424XG5cdFx0ZS0-PmQ6IOi9ieeZvFxuXHRcdGQtLT4-ZTog6L2J55m8XG4gICAgZS0tPj5jOiDlm57lgrPmlbjnorznjbjos4foqIpcbiAgICBsb29wIFNlcnZlci1TdHJlYW3kuLLmtYFcbiAgICAgIGMtPj5lOiBnUlBDLkRpZ2ltb24uUXVlcnlTdHJlYW08YnI-5Lul5pW456K85424SUTmkojlj5bmlbjnorznjbjos4foqIpcblx0XHRcdGUtPj5kOiDovYnnmbxcbiAgICAgIGxvb3Ag6ZuZ5ZCR5Liy5rWBXG4gICAgICAgIGQtPj53OiBnUlBDLldlYXRoZXIuUXVlcnk8YnI-5pKI5Y-W5q2k5pW456K85424J-S9jee9rkEn55qE5aSp5rCjXG4gICAgICAgIHctLT4-ZDog5Zue5YKz5aSp5rCjXG4gICAgICBlbmRcbiAgICAgIGQtLT4-ZTog5Zue5YKz5pW456K8542455qE6LOH6KiK6IiH5aSp5rCjXG5cdFx0XHRlLS0-PmM6IOi9ieeZvFxuICAgIGVuZFxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifX0)

可以發現流程圖與昨天的相似，不過多了 Envoy 來轉發。

## 先 run 起來再說

先 run 起來可能會比較有感覺，請 clone [Github-Example-Code](https://github.com/superj80820/2020-ithelp-contest)，並將 Server run 起來，

```
$ cd DAY15
$ docker-compose up
```

打開`localhost:8060`

![](https://i.imgur.com/GcOZ8ZP.png)

就可以看到 Web 端已經可以收到 Server-Stream 了！

## Web 講解

```javascript
const { QueryRequest, CreateRequest } = require("../proto/schema_pb.js");
const { DigimonPromiseClient } = require("../proto/schema_grpc_web_pb.js");

async function createDigimon(digimonPromiseClient, name) {
  let createRequest = new CreateRequest();
  createRequest.setName(name);
  const createResponse = await digimonPromiseClient.create(createRequest, {});
  return createResponse;
}

async function queryDigimonStream(digimonPromiseClient, digimonID) {
  let queryRequest = new QueryRequest();
  queryRequest.setId(digimonID);

  const queryStream = await digimonPromiseClient.queryStream(queryRequest, {});
  return queryStream;
}

(async () => {
  try {
    const digimonPromiseClient = new DigimonPromiseClient(
      "http://localhost:8080"
    );

    const createResponse = await createDigimon(digimonPromiseClient, "Agumon");
    const queryStream = await queryDigimonStream(
      digimonPromiseClient,
      createResponse.getId()
    );

    queryStream.on("data", function (response) {
      console.log(
        response.getId(),
        response.getName(),
        response.getStatus(),
        response.getLocation(),
        response.getWeather().toString()
      );
    });
    queryStream.on("status", function (status) {
      console.log(status.code);
      console.log(status.details);
      console.log(status.metadata);
    });
    queryStream.on("end", function (end) {
      // stream end signal
    });
  } catch (err) {
    console.error(err.code);
    console.error(err.message);
  }
})();
```

Web 的 Code 並不長，先以`DigimonPromiseClient()`建立 gRPC Client，之後就可以使用`queryStream`來獲得 stream 的物件。

stream 物件有一個 on function，可以以 callback 分別註冊:

- data: 接收到 Server-Stream 的資料
- status: 連線狀態
- end: 連線結束的處理

---

謝謝你的閱讀～
