本文章同時發佈於：

- [Github(包含程式碼)](https://github.com/superj80820/2020-ithelp-contest/blob/master/DAY12)
- [Medium](https://medium.com/%E9%AB%92%E6%A1%B6%E5%AD%90/day12-grpc-web-%E8%AE%93%E4%BD%A0%E7%9A%84%E5%89%8D%E7%AB%AF%E4%B9%9F%E5%90%83%E5%88%B0-grpc-%E7%9A%84%E6%83%A1%E9%AD%94%E6%9E%9C%E5%AF%A6-%E5%AF%A6%E4%BD%9C%E7%AF%87-95857f02595d)
- [iT 邦幫忙](https://ithelp.ithome.com.tw/articles/10244296)

```
文章為自己的經驗與夥伴整理的內容，設計沒有標準答案，如有可以改進的地方，請告訴我，我會盡我所能的修改，謝謝大家～
```

---

大家好，這次要介紹 gRPC-Web 的實作，大家可以直接 Clone 此份[Github 程式碼](https://github.com/superj80820/GRPC-Web-Simple-Demo)，直接先 run 起來會比較有感覺～

## 怎麼 run 起來？

1. `docker run -v ${PWD}:/app -w /app node:14.3.0-alpine npm install`: 安裝`web-cleint`的相依
2. `docker run -v ${PWD}:/app -w /app node:14.3.0-alpine npm run build`: 編譯`web-client`的`client.js`
3. `docker-compose build`: 安裝`Golang GRPC Server`與`Envoy`的相依
4. `docker-compose up`: 啟動`web-client`, `Golang GRPC Server`, `Envoy`
5. 連入`localhost:8060`，打開`console`會看到`web-cleint`發出的`Hello`得到`Hello World`的回傳

![](https://imgur.com/0GmF7QJ.jpg)

## envoy 設定

整體架構如 DAY11 所說的，後端與前端之間多了一個 envoy proxy，來把瀏覽器的 HTTP1.1 request 轉換成 HTTP2，

![](https://i.imgur.com/pJDltYR.png)

圖片來源: [Envoy and gRPC-Web: a fresh new alternative to REST](https://blog.envoyproxy.io/envoy-and-grpc-web-a-fresh-new-alternative-to-rest-6504ce7eb880)

所以，我們要在 docker-compose 設定前端、後端、envoy proxy 三者，

```yaml
# docker-compose.yml
version: "3"
services:
  envoy:
    build:
      context: ./
      dockerfile: ./docker/envoy/Dockerfile
    image: grpcweb/envoy
    ports:
      - "8080:8080"
      - "9901:9901"
    links:
      - server
  server:
    image: golang:1.14.6-alpine
    volumes:
      - ${PWD}:/server
    working_dir: /server
    ports:
      - "8070:8070"
    entrypoint: go run server/main.go
  web:
    image: httpd:2.4
    volumes:
      - ${PWD}/web-client:/usr/local/apache2/htdocs/
    ports:
      - "8060:80"
    links:
      - envoy
```

注意`dockerfile: ./docker/envoy/Dockerfile`處，我們必須 build 一個 envoy docker image，目標是把`envoy.yaml`放入 image 中，如下:

```yaml
# docker/envoy/Dockerfile
FROM envoyproxy/envoy:v1.15.0

COPY ./envoy.yaml /etc/envoy/envoy.yaml

CMD /usr/local/bin/envoy -c /etc/envoy/envoy.yaml -l trace --log-path /tmp/envoy_info.log
```

而 envoy 會依照此 config 來設定 proxy，

```yaml
# envoy.yaml
admin:
  access_log_path: /tmp/admin_access.log
  address:
    socket_address: { address: 0.0.0.0, port_value: 9901 }

static_resources:
  listeners:
    - name: listener_0
      address:
        # 講解1
        socket_address: { address: 0.0.0.0, port_value: 8080 }
      filter_chains:
        - filters:
            - name: envoy.filters.network.http_connection_manager
              typed_config:
                "@type": type.googleapis.com/envoy.config.filter.network.http_connection_manager.v2.HttpConnectionManager
                codec_type: auto
                stat_prefix: ingress_http
                route_config:
                  name: local_route
                  virtual_hosts:
                    - name: local_service
                      domains: ["*"]
                      # 講解2
                      routes:
                        - match: { prefix: "/" }
                          route:
                            cluster: echo_service
                            max_grpc_timeout: 0s
                      cors:
                        allow_origin_string_match:
                          - prefix: "*"
                        allow_methods: GET, PUT, DELETE, POST, OPTIONS
                        allow_headers: keep-alive,user-agent,cache-control,content-type,content-transfer-encoding,custom-header-1,x-accept-content-transfer-encoding,x-accept-response-streaming,x-user-agent,x-grpc-web,grpc-timeout
                        max_age: "1728000"
                        expose_headers: custom-header-1,grpc-status,grpc-message
                http_filters:
                  # 講解4
                  - name: envoy.filters.http.grpc_web
                  - name: envoy.filters.http.cors
                  - name: envoy.filters.http.router
  clusters:
    - name: echo_service
      connect_timeout: 0.25s
      type: logical_dns
      http2_protocol_options: {}
      lb_policy: round_robin
      load_assignment:
        cluster_name: cluster_0
        endpoints:
          - lb_endpoints:
              - endpoint:
                  address:
                    # 講解3
                    socket_address:
                      address: server
                      port_value: 50052
```

在`講解1`處，為 envoy 對外的 port 號，request 依照此 port 號進入後，可以依照`講解2`的 routes 來設定路由，如`"/"`導到了`clusters`裡的`echo_service`，而`echo_service`裡就可以實際將 requests 導到 docker-compose 中的 server，而 port 號是 50052。

而`講解4`處，因為有設定`envoy.filters.http.grpc_web`，所以對 HTTP1.1 流量會轉換成 HTTP2 的流量。

## 來產生 JS 與 Golanag 的 gRPC code

在 script 資料夾中，我已經寫好了 JS 與 Golang 產 gRPC code 的 script，分別如下：

```bash
# scirpt/build-proto-js.sh
#! /bin/bash

docker run \
    -v "${PWD}/proto:/protofile" \
    -e "protofile=helloworld.proto" \
    juanjodiaz/grpc-web-generator

mv ./proto/generated/* ./proto

rm -r proto/generated
```

```bash
# scirpt/build-proto-go.sh
docker run \
    -v "${PWD}/proto:/protofile" \
    -e "protofile=helloworld.proto" \
    juanjodiaz/grpc-web-generator

mv ./proto/generated/* ./proto

rm -r proto/generated
```

主要是透過 juanjodiaz 大大寫的[juanjodiaz/grpc-web-generator](https://github.com/juanjoDiaz/grpc-web-generator)來產生 gRPC code，省去了很多 gRPC 安裝的時間 XD。

run 完之後就會產生以下的 code，我們要 import 他們到我們實際的 code 中。

![](https://i.imgur.com/otM2XRI.png)

```javascript
// web-client/client.js

const { HelloRequest } = require("../proto/helloworld_pb.js");
const { GreeterPromiseClient } = require("../proto/helloworld_grpc_web_pb.js");

(async () => {
  try {
    const greeterService = new GreeterPromiseClient("http://localhost:8080");
    const request = new HelloRequest();
    request.setMessage("Hello");

    const response = await greeterService.sayHello(request, {});
    console.log(response.getMessage());
  } catch (err) {
    console.log(err.code);
    console.log(err.message);
  }
})();
```

前端中，我們利用`GreeterPromiseClient`來產生 gRPC client，並產生`HelloRequest`這個 message，使用`setMessage`將裡頭的 value 設定好，並透過`sayHello`傳送。

可以看到，`sayHello`使用`HelloRequest`是完全符合 gRPC 的 schema 的，前後端的定義都來自於 schema，溝通更 match 惹～

```proto
# proto/helloworld.proto
syntax = "proto3";

package helloworld;

service Greeter {
  rpc SayHello(HelloRequest) returns (HelloReply);
}

message HelloRequest {
  string message = 1;
}

message HelloReply {
  string message = 1;
}
```

---

謝謝你的閱讀～

## 參考

- [Envoy and gRPC-Web: a fresh new alternative to REST](https://blog.envoyproxy.io/envoy-and-grpc-web-a-fresh-new-alternative-to-rest-6504ce7eb880)
- [grpc-web-generator](https://github.com/juanjoDiaz/grpc-web-generator)
