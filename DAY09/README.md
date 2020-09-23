本文章同時發佈於：

- [Github(包含程式碼)](https://github.com/superj80820/2020-ithelp-contest/blob/master/DAY09)
- [Medium](https://medium.com/%E9%AB%92%E6%A1%B6%E5%AD%90/day9-%E7%A5%9E%E5%A5%87%E7%9A%84-grpc-%E8%AE%93%E4%BD%A0%E6%8A%8A-call-service-%E7%95%B6%E6%88%90%E4%B8%80%E5%80%8B-function-call-%E6%A6%82%E5%BF%B5%E7%AF%87-1a27e7f24331)
- [iT 邦幫忙](https://ithelp.ithome.com.tw/articles/10242372)

```
文章為自己的經驗與夥伴整理的內容，設計沒有標準答案，如有可以改進的地方，請告訴我，我會盡我所能的修改，謝謝大家～
```

---

大家好，繼昨天[DAY09](https://github.com/superj80820/2020-ithelp-contest/blob/master/DAY07)的介紹後，Clean Architecture 架構告一段落。現在我們來思考一下，如果後端微服務龐大起來，如下圖，可能會產生什麼問題？

## 微服務溝通的災難與解決

![](https://i.imgur.com/M4NGgqF.png)

圖片來源: [微服務架構之「 下一代微服務 Service Mesh 」](https://www.itread01.com/content/1561532706.html)

- 傳輸的速度因為 JSON[文本傳輸](https://zh.wikipedia.org/wiki/%E8%B6%85%E6%96%87%E6%9C%AC%E4%BC%A0%E8%BE%93%E5%8D%8F%E8%AE%AE)的方式導致緩慢
- 大量的服務溝通會使得 Restful API 的 request 與 response 的欄位逐漸混亂

雖然我們已經透過了 Swagger 最大程度的控管 Restful API 的介面，但`值透過enum來實現`、`這個參數是require但你忘記帶程式也不會告知你`，為了解決這些問題，有個很簡單粗暴的方法：

> 讓 Service 的溝通變成 function call 吧！

而這件事可以透過[RPC(Remote Procedure Call)](https://zh.wikipedia.org/zh-tw/%E9%81%A0%E7%A8%8B%E9%81%8E%E7%A8%8B%E8%AA%BF%E7%94%A8)達成，流程圖如下，其實他得概念並不複雜，單純就是把`call function與return result間的流程加了一道網路轉換`

![](https://i.imgur.com/i51FX43.png)

此圖來源: [Day 16 - 分散式系統溝通的方法 - RPC](https://ithelp.ithome.com.tw/articles/10223580)

1. Client 端呼叫一個 function
2. 此 function 的參數會被編碼成特定格式的封包，供網路傳送使用
3. 網路傳送封包
4. Server 接收封包
5. 將封包轉換成 function 的參數格式
6. 執行此 function，並回傳 result

事實上 RPC 在 1976 年就誕生了，並不是什麼新技術，當時為了要處理多台電腦的通信採用此方法，後來衍伸出了[XML-RPC](https://zh.wikipedia.org/wiki/XML-RPC)、[SOAP](https://zh.wikipedia.org/wiki/%E7%AE%80%E5%8D%95%E5%AF%B9%E8%B1%A1%E8%AE%BF%E9%97%AE%E5%8D%8F%E8%AE%AE)這樣的協定，但在定義與使用上都有些繁瑣古老，後來 Google 發展出了[gRPC](https://zh.wikipedia.org/wiki/GRPC)，為現代的微服務提供簡單高效強大的 RPC 協定。

## 新世代的 RPC - gRPC

![](https://i.imgur.com/PoiSXeG.png)

- 傳輸緩慢的問題: gPRC 透過[Protobuf](https://zh.wikipedia.org/wiki/Protocol_Buffers)轉換來解決，會將 function 的參數轉換成`二進制`的格式，以支援二進制溝通功能的[HTTP/2](https://zh.wikipedia.org/wiki/HTTP/2#HTTP/2%E4%B8%8EHTTP/1.1%E6%AF%94%E8%BE%83)來傳送，使得溝通極小極快。
- 微服務溝通介面混亂的問題: gRPC 透過定義 Protobuf 的 schema 解決，此 schema 不但可以定義傳輸的訊息，也可以定有哪些傳輸 function，並且這些 function 的參數是什麼。

以 Golang 來說，這些定好的 Protobuf schema 透過官方的[codegen](https://www.grpc.io/docs/languages/go/quickstart/) tool，直接產生實際的 function 與 struct。

而你實際使用時，只要直接 call 此 function 並且帶對參數就可以了，因為是 function call，所以要帶什麼參數更加的直觀，而編譯器會告訴你是否帶對參數也更佳安全。

### 多種便利的傳輸方案

gRPC 在定義 function 時，可以針對不同場景，使用多種方案:

- 單向傳輸(Unary): 機制上跟 Restful API 一樣，就是一個 request 一個 response，只不過是透過了 gRPC 協定。
- 單向串流(Streaming): 你可以想像成訂閱機制，client 透過`一個 request`訂閱 server 後，server 就可以發送`無數次 response`。
- 雙向串流(Bidirectional streaming): 機制就像 Websocket，但是在傳輸格式有透過 schema 定義，client 可以發送無數次 request，server 也可以發送無數次 response

## 別忘了 RPC 還是透過網路

微服務很大的痛就是網路，試想看看 function call 有可能會因為網路不穩而失敗，這會多難麻煩 XD。

gRPC 可用此方法處理，Deadlines/Timeouts 來設定 timeout，[go-grpc-middleware](https://godoc.org/github.com/grpc-ecosystem/go-grpc-middleware/retry)來 retry。

總而言之這是個要特別注意的地方，雖然變成了 function call 但他還是`網路傳輸`XD。

---

接下來將介紹 gRPC 的傳輸方案的實際實作，謝謝你的閱讀～

## 參考

- [浅谈 RPC 和 REST: SOAP, gRPC, REST - 知乎](https://zhuanlan.zhihu.com/p/60352360)
- [web-SOA 和微服務之間有什麼區別？](https://stackoverflow.com/questions/48190148/what-is-the-difference-between-soa-and-microservices)
- [簡單物件存取協定 - 維基百科，自由的百科全書](https://zh.wikipedia.org/wiki/%E7%AE%80%E5%8D%95%E5%AF%B9%E8%B1%A1%E8%AE%BF%E9%97%AE%E5%8D%8F%E8%AE%AE)
- [XML-RPC - 維基百科，自由的百科全書](https://zh.wikipedia.org/wiki/XML-RPC)
- [微服務架構之「 下一代微服務 Service Mesh 」](https://www.itread01.com/content/1561532706.html)
- [超文本傳輸協定 - 維基百科，自由的百科全書](https://zh.wikipedia.org/wiki/%E8%B6%85%E6%96%87%E6%9C%AC%E4%BC%A0%E8%BE%93%E5%8D%8F%E8%AE%AE)
- [Day 16 - 分散式系統溝通的方法 - RPC - iT 邦幫忙::一起幫忙解決難題，拯救 IT 人的一天](https://ithelp.ithome.com.tw/articles/10223580)
- [遠端程序呼叫 - 維基百科，自由的百科全書](https://zh.wikipedia.org/wiki/%E9%81%A0%E7%A8%8B%E9%81%8E%E7%A8%8B%E8%AA%BF%E7%94%A8)
- [gRPC - 維基百科，自由的百科全書](https://zh.wikipedia.org/wiki/GRPC)
- [Protocol Buffers - 維基百科，自由的百科全書](https://zh.wikipedia.org/wiki/Protocol_Buffers)
- [HTTP/2 - 維基百科，自由的百科全書](https://zh.wikipedia.org/wiki/HTTP/2#HTTP/2%E4%B8%8EHTTP/1.1%E6%AF%94%E8%BE%83)
- [Core concepts, architecture and lifecycle – gRPC](https://grpc.io/docs/what-is-grpc/core-concepts/)
