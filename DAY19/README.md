本文章同時發佈於：

- [Github(包含程式碼)](https://github.com/superj80820/2020-ithelp-contest/blob/master/DAY19)
- [Medium](https://medium.com/%E9%AB%92%E6%A1%B6%E5%AD%90/day19-%E4%BA%86%E8%A7%A3-k8s-%E4%B8%AD%E7%9A%84%E5%A4%A7%E9%96%80%E7%A5%9E-ingress-ee996d8837ae)
- [iT 邦幫忙](https://ithelp.ithome.com.tw/articles/10248658)

```
文章為自己的經驗與夥伴整理的內容，設計沒有標準答案，如有可以改進的地方，請告訴我，我會盡我所能的修改，謝謝大家～
```

---

大家好，今天我們要介紹 K8s 的另一個重要元件`Ingress`，Ingress 是一個在 K8s 眾多 Service 前的 Reverse Proxy，他可以讓 K8s 的連線變為:

> 一個 domain，由不同的 hostName 或 pathName 來連線

![](https://i.imgur.com/AVb15MY.png)

圖片來源: [Kubernetes 基礎教學（二）實作範例：Pod、Service、Deployment、Ingress](https://medium.com/@C.W.Hu/kubernetes-implement-ingress-deployment-tutorial-7431c5f96c3e)

如圖舉個例子，如果我們只有`digimon.com`一個 domain，卻有 a, b, c, d, e 5 個 service 怎麼辦？

- hostName 的分流做法: 設置子網域 a.digimon.com, b.digimon.com ... 等等
- pathName 的分流做法: 設置路徑 digimon.com/a, digimon.com/b ... 等等

這樣就不用另外再去申請 domain 了！(~~當然如果你很有錢也可以啦 XD~~)

## 實作

Ingress 有許多不同環境的版本，ingress-gce, ingress-aws, ingress-nginx 等等(詳細在[additional-controllers](https://kubernetes.io/docs/concepts/services-networking/ingress-controllers/#additional-controllers))，如果是使用 ingress-aws，他在實際套用時會直接與 AWS 的 Load Balancer 做連動，而我們在 local 直接使用預設的即可。

以下是以 pathName 來做分流，

```yaml
apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  name: digimon-ingress
spec:
  rules:
    - host: digimon.com
      http:
        paths:
          - path: /digimon-service
            backend:
              serviceName: envoy
              servicePort: 8080
```

- rules.host: 設置實際上連線的 domain
- rules.http.paths.path: 設置分流 path，在此會以`digimon.com/digimon-service`來導到正確的 envoy-service
- rules.http.paths.backend.serviceName: 要導到的 service name
- rules.http.paths.backend.servicePort: 要導到的 service port

由於我使用上發現 ingress 要啟用要使用 hyperkit，所以可以先刪除 minikube 在重新建立，

```bash
$ minikube delete
$ minikube start --driver=hyperkit --kubernetes-version v1.15.0
```

啟用 ingress

```bash
minikube addons enable ingress
```

啟動 service, deployment, ingress

```
$ kompose up
$ kubectl create -f ingress.yaml
```

獲取 minikube ip

```
$ minikube ip
```

由於 domain 並不真的存在，所以我們可以以修改`/etc/hosts`的方式，來將 domain 導到 minikube ip

我得到的 minikube ip 是`127.0.0.1`，因此我可以下此指令將 domain.com 導到 127.0.0.1

```
echo 127.0.0.1   digimon.com  >> /etc/hosts
```

[//]: #"yorktodo將前後端分離，改固定ip"

[//]: #"yorktodo envoy 要在最後啟動這個問題需要解決"
[//]: #"yorktodo`kubectl create -f db-deployment.yaml,db-service.yaml,envoy-deployment.yaml,envoy-service.yaml,server-deployment.yaml,server-service.yaml,weather-deployment.yaml,weather-service.yaml,web-deployment.yaml,web-service.yaml`"
[//]: #"yorktodo 修改文章 etc/host 為 192.168.64.7 api.backend.com"

## 參考

- [Kubernetes 基礎教學（二）實作範例：Pod、Service、Deployment、Ingress | by 胡程維｜ Cheng-Wei Hu | Medium](https://medium.com/@C.W.Hu/kubernetes-implement-ingress-deployment-tutorial-7431c5f96c3e)
- [将 Docker Compose 文件转换为 Kubernetes 资源 | Kubernetes](https://kubernetes.io/zh/docs/tasks/configure-pod-container/translate-compose-kubernetes/)
- [Ingress path does not work · Issue #349 · kubernetes/ingress-nginx](https://github.com/kubernetes/ingress-nginx/issues/349)
- [Set up Ingress on Minikube with the NGINX Ingress Controller | Kubernetes](https://kubernetes.io/docs/tasks/access-application-cluster/ingress-minikube/)
- [docker: Ingress not exposed on MacOS · Issue #7332 · kubernetes/minikube](https://github.com/kubernetes/minikube/issues/7332)
- [Ingress Controllers | Kubernetes](https://kubernetes.io/docs/concepts/services-networking/ingress-controllers/#additional-controllers)
