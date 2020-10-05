本文章同時發佈於：

- [Github(包含程式碼)](https://github.com/superj80820/2020-ithelp-contest/blob/master/DAY21)
- [Medium](https://medium.com/%E9%AB%92%E6%A1%B6%E5%AD%90/day21-istio-%E6%98%AF%E4%BB%80%E9%BA%BC-%E5%8E%9F%E4%BE%86%E6%98%AF%E5%AE%B9%E5%99%A8%E9%96%93%E9%80%9A%E8%A8%8A%E7%9A%84%E5%A6%BB%E7%AE%A1%E5%9A%B4%E5%97%8E-4fe0b9c0a1d8)
- [iT 邦幫忙](https://ithelp.ithome.com.tw/articles/10249132)

```
文章為自己的經驗與夥伴整理的內容，設計沒有標準答案，如有可以改進的地方，請告訴我，我會盡我所能的修改，謝謝大家～
```

---

大家好，在 K8s 告一段落後，接下來要來介紹 Istio，

![](https://i.imgur.com/qKTWiW3.png)

圖片來源: [gRPC in Microservices](https://levelup.gitconnected.com/grpc-in-microservices-5887caef195)

K8s 與 Istio 都是平台層的架構，而他們的職責分別是:

- K8s: 管理容器生命週期
- Istio: 管理容器間的通訊

你可能會想`K8s`不是也能透過 load balance 來管理流量之類的通訊嗎？

是的，K8s 與 Istio 有著許多相似的功能，但是要注意的是 K8s 從頭到尾都是為了`管理容器生命週期`而生，`管理容器間的通訊`並不是他的強項。

## 容器間的通訊

在以往，容器間的通訊，往往是透過直接在容器裡的 code 來實作，

比如說 A service 呼叫 B service 失敗了，那就需要 retry，而 retry 就在 service 的 code 中實作，

你可以選擇在 code 中實作，但勢必就會造成 code 中除了有許多`應用層的邏輯`以外，又多了許多`平台層的邏輯`，這使得 microservice 變得很不 microservice，因為添加了許多

> 非此 microservice 邏輯的 code

或者你使用[SOA(service-oriented architecture)](https://zh.wikipedia.org/wiki/%E9%9D%A2%E5%90%91%E6%9C%8D%E5%8A%A1%E7%9A%84%E4%BD%93%E7%B3%BB%E7%BB%93%E6%9E%84)框架例如 Spring Cloud、Go-Micro，這樣的框架幫你處理完了容器間通訊的問題，但這類框架通常與 K8s 整合上會遇到些問題，你可以參考[當 Spring Cloud 遇上 Kubernetes-完整的微服務框架從此不完整了](https://medium.com/brobridge/%E7%95%B6-spring-cloud-%E9%81%87%E4%B8%8A-kubernetes-5bc9e6ce602f)，最根本的原因就是

> 框架功能與 K8s 平台層的功能重疊

Spring Cloud 來說，基本上就可以當作一個完整的微服務平台了，其實不需要 K8s，但你的 backend 就無法與主流 K8s 靠攏，如果硬靠又會很多的見招拆招。

## SOA 與 microservice 的差異

[SOA(service-oriented architecture)](https://zh.wikipedia.org/wiki/%E9%9D%A2%E5%90%91%E6%9C%8D%E5%8A%A1%E7%9A%84%E4%BD%93%E7%B3%BB%E7%BB%93%E6%9E%84)是比 microserivce 還早出現的一個名詞，中文為`服務導向架構`，在精神上，microservice 相同都是`將服務切分為更小的服務組合`，但如圖，

![](https://i.imgur.com/4JXCBNA.png)

圖片來源: [What is the difference between SOA and Microservices](https://stackoverflow.com/questions/48190148/what-is-the-difference-between-soa-and-microservices)

SOA 更注重在一個框架中完成所以服務的組合，彼此間的呼叫還是被此框架掌控，

而現在的 microservice 專注於將服務完全切分，服務與服務的呼叫是沒有依靠任何框架或插件。

## 服務溝通既沒有直接依賴框架或插件，又可以管理的聰明辦法 - Istio Sidecar

統整上述有幾個問題需要實現:

1. 在不修改容器的前提下實現容器通訊的管理
2. 容器通訊不直接依賴特定框架或插件
3. 以 K8s 為基底設計

而 Istio 想到的方法就是 Sidecar，概念很簡單:

> 在每個 microserivce 旁新增一個 Envoy Proxy 來劫持所有流量，並統一控制所有 proxy 來管理通訊

![](https://i.imgur.com/dFVaNT0.png)

如圖，圖中所有 Envoy Proxy 的部分都是附屬在 microservice 上的，如果我們想要加入 Istio，對我們整體的 K8s 容器是沒有太大影響的，因為他並不是實作在容器中，而是向妻管嚴這樣從旁拿取所有流量。

## 結論

Istio 管理這些容器通訊可以稱為`Service Mesh`，有以下幾點有用的功能:

- 請求轉發：如服務發現，負載平衡
- 流量控制：更加彈性的分流
- 服務發布：使用金絲雀，A/B 等等方式來部署
- 錯誤管理：現流，熔斷，重試，
- 安全性：驗證，授權，加密

---

並且 Istio 可以讓一切的容器管理更加`平台層化`，這樣後端 RD 可以更專注業務邏輯開發，而架構師可以更專注平台維護。

## 參考

- [蚂蚁金服 Service Mesh 实践探索 - 知乎](https://zhuanlan.zhihu.com/p/48105816)
- [導入微服務前一定要知道的事 | iThome](https://www.ithome.com.tw/news/116053)
- [web - What is the difference between SOA and Microservices - Stack Overflow](https://stackoverflow.com/questions/48190148/what-is-the-difference-between-soa-and-microservices)
- [gRPC in Microservices. Originally published here. | by Milad Irannejad | Level Up Coding](https://levelup.gitconnected.com/grpc-in-microservices-5887caef195)
- [服務導向架構 - 維基百科，自由的百科全書](https://zh.wikipedia.org/wiki/%E9%9D%A2%E5%90%91%E6%9C%8D%E5%8A%A1%E7%9A%84%E4%BD%93%E7%B3%BB%E7%BB%93%E6%9E%84)
- [當 Spring Cloud 遇上 Kubernetes. 完整的微服務框架從此不完整了 | by Fred Chien（錢逢祥） | Brobridge - 寬橋微服務 | Medium](https://medium.com/brobridge/%E7%95%B6-spring-cloud-%E9%81%87%E4%B8%8A-kubernetes-5bc9e6ce602f)
- [Istio / 在虚拟机上部署 Bookinfo 应用程序](https://istio.io/latest/zh/docs/examples/virtual-machines/bookinfo/)
