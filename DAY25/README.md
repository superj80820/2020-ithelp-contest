本文章同時發佈於：

- [Github(包含程式碼)](https://github.com/superj80820/2020-ithelp-contest/blob/master/DAY25)

```
文章為自己的經驗與夥伴整理的內容，設計沒有標準答案，如有可以改進的地方，請告訴我，我會盡我所能的修改，謝謝大家～
```

---

大家好，接下來要介紹Istio DestinationRule元件，他是用來處理流量導流的部分，可以靠他來處理Load Balancer，你可能想問`為什麼不直接用K8s的Load Balancer`呢？以我目前最大的感受是

> K8s Load Balancer無法分配HTTP2的長連線

## GRPC的特性產生的問題

傳統的Restful API，都是一個request一個response，而GRPC連線是一個長連線，並且利用此連線來連續送request與連續接收response，

傳統的Restful API可以為每個requests做分流:

![](https://i.imgur.com/eWHvyf5.png)

[//]: #"(./restfulapi.drawio.png)"

而GRPC的長連線會導致分流失敗:

![](https://i.imgur.com/56uwaCZ.png)

[//]: #"(./grpc.drawio.png)"

## 使用一個能解析的HTTP的Proxy來解決問題

要解決長連線分配請求的方式很直覺，就是

> 實際解析長連線裡頭的內容

所以我們必須要有一個Proxy能夠看懂HTTP傳送的內容，不能再單純是轉發requests的Load Balancer了，

而Envoy就可以當作此請求的Proxy，他解析之後就會如下圖發送請求:

![](https://i.imgur.com/VOFPBs0.png)

[//]: #"(./istiogrpc.drawio.png)"

## 參考

* [Istio / Destination Rule](https://istio.io/latest/docs/reference/config/networking/destination-rule/)
* [Istio 基礎 — gRPC 負載均衡. 什麼是 Istio | by Alan Chen | getamis | Medium](https://medium.com/getamis/istio-%E5%9F%BA%E7%A4%8E-grpc-%E8%B2%A0%E8%BC%89%E5%9D%87%E8%A1%A1-d4be0d49ee07)
* [[Day03]為什麼我們要用Istio，Native Kubernetes有什麼做不到 - iT 邦幫忙::一起幫忙解決難題，拯救 IT 人的一天](https://ithelp.ithome.com.tw/articles/10217407)
* [在不流淚的Kubernetes上實現gRPC負載平衡 Kubernetes](https://kubernetes.io/blog/2018/11/07/grpc-load-balancing-on-kubernetes-without-tears/)