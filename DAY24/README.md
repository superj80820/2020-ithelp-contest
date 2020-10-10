本文章同時發佈於：

- [Github(包含程式碼)](https://github.com/superj80820/2020-ithelp-contest/blob/master/DAY24)
- [Medium]()
- [iT 邦幫忙]()

```
文章為自己的經驗與夥伴整理的內容，設計沒有標準答案，如有可以改進的地方，請告訴我，我會盡我所能的修改，謝謝大家～
```

---

大家好，今天要介紹的是Istio Gateway、VirtualService的概念，`為何不直接使用K8s Ingress元件？`，`為什麼要分那麼多層元件呢？`

## K8s Ingress無法達到Service Mesh的功能

[官方說明](https://istio.io/latest/blog/2018/v1alpha3-routing/):

> For ingress traffic management, you might ask: Why not reuse Kubernetes Ingress APIs? The Ingress APIs proved to be incapable of expressing Istio’s routing needs. By trying to draw a common denominator across different HTTP proxies, the Ingress is only able to support the most basic HTTP routing and ends up pushing every other feature of modern proxies into non-portable annotations.

簡單來說，在原本的K8s Ingress元件，無法提供Service Mesh的需求，僅適用簡單的路由，而要有辦法精確的控制Service Mesh，必須要將Ingress的職責切分的更細，如以下:

```
外部Load Balancer --> Istio Gateway元件 --> Istio VirtualService元件 --> K8s Service元件 --> K8s Pods元件
```

官方提供了下圖:

![](https://i.imgur.com/aSzOvHl.png)

Istio將L4-L6層的功能(如TCP、TSL等等)，交給Gateway元件處理，

而L7層的應用層就可以在特別客製，如果要連接不同的Service就交給VirtualService元件處理，如果要管理流量就透過DestinationRule元件處理，

## 結論

在K8s搭配Istio的系統中，K8s許多溝通的功能會盡量讓Istio來做，系統會越來越趨向於`Istio負責各種溝通、K8s只負責容器的生命週期`，所以你會發現Istio慢慢將K8s的溝通細分掌控，這樣才能達到Service Mesh的好處。

## 參考

* [TCP vs HTTP(S) Load Balancing.. In this article I shall show two main… | by Martin Ombura Jr. | Martin Ombura Jr. | Medium](https://medium.com/martinomburajr/distributed-computing-tcp-vs-http-s-load-balancing-7b3e9efc6167)
* [Istio / Introducing the Istio v1alpha3 routing API](https://istio.io/latest/blog/2018/v1alpha3-routing/)
* [[Day17] 如何為Cluster選擇一個好的Gateway - iT 邦幫忙::一起幫忙解決難題，拯救 IT 人的一天](https://ithelp.ithome.com.tw/articles/10224374)
* [流量管理基础概念 · Istio Handbook - Istio 服务网格进阶实战 by Jimmy Song(宋净超)](https://jimmysong.io/istio-handbook/concepts/traffic-management-basic.html)