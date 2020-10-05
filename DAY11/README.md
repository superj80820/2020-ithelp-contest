本文章同時發佈於：

- [Github(包含程式碼)](https://github.com/superj80820/2020-ithelp-contest/blob/master/DAY11)
- [Medium](https://medium.com/%E9%AB%92%E6%A1%B6%E5%AD%90/day11-grpc-web-%E8%AE%93%E4%BD%A0%E7%9A%84%E5%89%8D%E7%AB%AF%E4%B9%9F%E5%90%83%E5%88%B0-grpc-%E7%9A%84%E6%83%A1%E9%AD%94%E6%9E%9C%E5%AF%A6-%E6%A6%82%E5%BF%B5%E7%AF%87-cc001c9fec5b)
- [iT 邦幫忙](https://ithelp.ithome.com.tw/articles/10243651)

```
文章為自己的經驗與夥伴整理的內容，設計沒有標準答案，如有可以改進的地方，請告訴我，我會盡我所能的修改，謝謝大家～
```

---

大家好，在昨天的 gRPC 實作後，在微服務上我們可以用更完整的方案溝通，那是否有一個方案，可以將 gRPC 運用在前端與後端溝通呢？

有的，那就是[gRPC-Web](https://github.com/grpc/grpc-web)！

## 為什麼需要 gRPC-Web

![](https://i.imgur.com/ZWpEW7h.png)

圖片來源: [Envoy and gRPC-Web: a fresh new alternative to REST](https://blog.envoyproxy.io/envoy-and-grpc-web-a-fresh-new-alternative-to-rest-6504ce7eb880)

gRPC 在前端運作最大的問題，就是前端對於 gRPC 所使用的 protocol 不是那麼完好的支援，現在主流瀏覽器使用`HTTP1.1`，而 gRPC 使用的是更加高效的`HTTP2`。

那該怎麼解決，Google 提供的解決方案是利用一個 proxy 架設在後端，將前端所有的 HTTP1.1 流量轉換成 HTTP2，如下圖:

![](https://i.imgur.com/pJDltYR.png)

圖片來源: [Envoy and gRPC-Web: a fresh new alternative to REST](https://blog.envoyproxy.io/envoy-and-grpc-web-a-fresh-new-alternative-to-rest-6504ce7eb880)

envoy 是現在常見的一個 proxy，他對於微服務溝通的各種問題有很好的緩解作用，在未來 istio 介紹中，會更常看到他。

有這個 HTTP1.1 透過 envoy 轉換的過程，勢必會缺小原生 HTTP2 的一些好處，如下:

- 不支援 Client-Streaming: 雖然支援 Server-Streaming 的功能，但在 Client-Streaming 上因為 HTTP1.1 的[安全性問題](https://github.com/grpc/grpc-web/blob/master/doc/streaming-roadmap.md#client-streaming-and-half-duplex-streaming)，始終沒有啟用，一旦解決後 gRPC-Web 可以透過 Server 與 Client 的 Streaming 來做雙工串流。
- [不使用二進制傳輸](https://github.com/grpc/grpc-web/blob/master/doc/roadmap.md#non-binary-message-encoding): 因為前端使用二進制並非會有較好的效能，所以還是使用類似 JSON 的方式傳送。

所以，gPRC-Web 在選用上，如果你有以下需求，那他將會是你的好選擇:

- 前後端要有一致的溝通介面
- 前端有訂閱需求(可用 Server-Steaming 實作)

---

接下來會對 gRPC-Web 實際實作，謝謝你的閱讀～

## 參考

- [grpc / grpc-web：適用於 Web 客戶端的 gRPC](https://github.com/grpc/grpc-web)
- [master 上的 grpc-web / roadmap.md·grpc / grpc-web](https://github.com/grpc/grpc-web/blob/master/doc/roadmap.md#non-binary-message-encoding)
- [Envoy and gRPC-Web: a fresh new alternative to REST | by Luc Perkins | Envoy Proxy](https://blog.envoyproxy.io/envoy-and-grpc-web-a-fresh-new-alternative-to-rest-6504ce7eb880)
- [grpc-web/streaming-roadmap.md at master · grpc/grpc-web](https://github.com/grpc/grpc-web/blob/master/doc/streaming-roadmap.md#client-streaming-and-half-duplex-streaming)
