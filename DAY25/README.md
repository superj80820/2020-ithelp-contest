本文章同時發佈於：

- [Github(包含程式碼)](https://github.com/superj80820/2020-ithelp-contest/blob/master/DAY25)

```
文章為自己的經驗與夥伴整理的內容，設計沒有標準答案，如有可以改進的地方，請告訴我，我會盡我所能的修改，謝謝大家～
```

---

大家好，接下來要介紹 Istio DestinationRule 元件，他是用來處理流量導流的部分，可以靠他來處理 Load Balancer，你可能想問`為什麼不直接用K8s的Load Balancer`呢？以我目前最大的感受是

> K8s Load Balancer 無法分配 HTTP2 的長連線

## GRPC 的特性產生的問題

傳統的 Restful API，都是一個 request 一個 response，而 GRPC 連線是一個長連線，並且利用此連線來連續送 request 與連續接收 response，

傳統的 Restful API 可以為每個 requests 做分流:

![](https://i.imgur.com/eWHvyf5.png)

[//]: #"(./restfulapi.drawio.png)"

而 GRPC 的長連線會導致分流失敗:

![](https://i.imgur.com/56uwaCZ.png)

[//]: #"(./grpc.drawio.png)"

## 使用一個能解析的 HTTP 的 Proxy 來解決問題

要解決長連線分配請求的方式很直覺，就是

> 實際解析長連線裡頭的內容

所以我們必須要有一個 Proxy 能夠看懂 HTTP 傳送的內容，不能再單純是轉發 requests 的 Load Balancer 了，

而 Envoy 就可以當作此請求的 Proxy，他解析之後就會如下圖發送請求:

![](https://i.imgur.com/VOFPBs0.png)

[//]: #"(./istiogrpc.drawio.png)"

## 參考

- [Istio / Destination Rule](https://istio.io/latest/docs/reference/config/networking/destination-rule/)
- [Istio 基礎 — gRPC 負載均衡. 什麼是 Istio | by Alan Chen | getamis | Medium](https://medium.com/getamis/istio-%E5%9F%BA%E7%A4%8E-grpc-%E8%B2%A0%E8%BC%89%E5%9D%87%E8%A1%A1-d4be0d49ee07)
- [[Day03]為什麼我們要用 Istio，Native Kubernetes 有什麼做不到 - iT 邦幫忙::一起幫忙解決難題，拯救 IT 人的一天](https://ithelp.ithome.com.tw/articles/10217407)
- [在不流淚的 Kubernetes 上實現 gRPC 負載平衡 Kubernetes](https://kubernetes.io/blog/2018/11/07/grpc-load-balancing-on-kubernetes-without-tears/)
