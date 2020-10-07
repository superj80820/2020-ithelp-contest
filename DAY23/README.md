本文章同時發佈於：

- [Github(包含程式碼)](https://github.com/superj80820/2020-ithelp-contest/blob/master/DAY23)
- [Medium]()
- [iT 邦幫忙](https://ithelp.ithome.com.tw/articles/10250134)

```
文章為自己的經驗與夥伴整理的內容，設計沒有標準答案，如有可以改進的地方，請告訴我，我會盡我所能的修改，謝謝大家～
```

---

大家好。今天要介紹 Istio 的`VirtualService`、`Gateway`，兩者搭配可以達到 K8s 的`Ingress`的效果，並且有著 Istio 更強大的功能。

而這篇文章會以實作為介紹，實際原理將在之後介紹，

## 實作

以下的實作範例都在[Code-Example](https://github.com/superj80820/2020-ithelp-contest)中，

在`helm-digimon/templates/server-service.yaml`中，我們須將 ports 的name做修改為`grpc-web`，Istio在看到此name後，就會調整envoy proxy使client網頁端的`grpc-web`請求可成功與server溝通。[詳細API在此](https://istio.io/latest/docs/ops/configuration/traffic-management/protocol-selection/)

![](https://i.imgur.com/AiUJiVt.png)

新增`VirtualService`、`Gateway`兩個元件至`DAY23/helm-digimon/templates/gateway.yaml`

```yaml
apiVersion: networking.istio.io/v1alpha3
kind: Gateway
metadata:
  name: gateway
spec:
  selector:
    istio: ingressgateway
  servers:
    - port:
        number: 80
        name: http
        protocol: HTTP
      hosts:
        - "*.backend.com"
---
apiVersion: networking.istio.io/v1alpha3
kind: VirtualService
metadata:
  name: virtual-service
spec:
  hosts:
    - "api.backend.com"
  gateways:
    - gateway
  http:
    - match:
        - uri:
            prefix: /digimon.Digimon
      route:
        - destination:
            port:
              number: 6000
            host: server
      corsPolicy:
        allowOrigins:
          - exact: "http://web.backend.com"
        allowMethods:
          - POST
          - GET
          - OPTIONS
          - PUT
          - DELETE
        allowHeaders:
          - grpc-timeout
          - content-type
          - keep-alive
          - user-agent
          - cache-control
          - content-type
          - content-transfer-encoding
          - x-accept-content-transfer-encoding
          - x-accept-response-streaming
          - x-user-agent
          - x-grpc-web
        maxAge: 1728s
        exposeHeaders:
          - grpc-status
          - grpc-message
        allowCredentials: true
---
apiVersion: networking.istio.io/v1alpha3
kind: VirtualService
metadata:
  name: web
spec:
  hosts:
    - "web.backend.com"
  gateways:
    - gateway
  http:
    - match:
        - uri:
            prefix: /
      rewrite:
        uri: /
      route:
        - destination:
            port:
              number: 8060
            host: web
```

`DAY23/helm-digimon/templates/gateway.yaml`有一個Gateway、一個web的VirtualService、一個server的VirtualService，

外部的流量會如下導至正確的pods，

```
Istio Gateway元件 --> Istio VirtualService元件 --> K8s Service元件 --> K8s Pods元件
```

將 K8s run 起來，

```bash
$ minikube start --kubernetes-version v1.16.0
```

安裝Istio系統至K8s，並啟用envoy sidecar，

```bash
$ istioctl install --set profile=demo
$ kubectl label namespace default istio-injection=enabled
```

啟動所有K8s元件

```bash
$ cd DAY23/helm-digimon
$ helm install . --generate-name
```

Istio Gateway設計的外部流量導向如下:

```
外部Load Balancer --> Istio Gateway元件
```

由於實作是採用minikube，我們必須模擬外部Load Balancer出來，

```bash
$ minikube tunnel
```

模擬完畢後透過`minikube ip`或的minikube的外部ip，並透過`$ open /etc/hosts`將URL導至此ip，如果minikube ip是`127.0.0.1`的話修改方式如下:

![](https://i.imgur.com/oWYdLXm.png)

都完成後打開`web.backend.com`，web client已經成功與server溝通，太好了。

![](https://i.imgur.com/dl68jd6.png)


## 結論

你可能會想`為什麼要特別設計一個VirtualService元件，不直接用K8s Service元件呢？`，`為什麼要分那麼多層元件呢？`，以下問題都會在下篇文章做探討，謝謝你的閱讀～

## 參考

- [Istio / Getting Started](https://istio.io/latest/docs/setup/getting-started/)
- [Istio / What is Istio?](https://istio.io/latest/docs/concepts/what-is-istio/)
- [Deprecated APIs Removed In 1.16: Here’s What You Need To Know | Kubernetes](https://kubernetes.io/blog/2019/07/18/api-deprecations-in-1-16/)
