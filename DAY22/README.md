本文章同時發佈於：

- [Github(包含程式碼)](https://github.com/superj80820/2020-ithelp-contest/blob/master/DAY22)
- [Medium](https://medium.com/%E9%AB%92%E6%A1%B6%E5%AD%90/day22-istio-%E6%98%AF%E4%BB%80%E9%BA%BC-%E5%8E%9F%E4%BE%86%E6%98%AF%E5%AE%B9%E5%99%A8%E9%96%93%E9%80%9A%E8%A8%8A%E7%9A%84%E5%A6%BB%E7%AE%A1%E5%9A%B4%E5%97%8E-%E5%AF%A6%E4%BD%9C%E7%AF%87-3230f841c9ae)
- [iT 邦幫忙](https://ithelp.ithome.com.tw/articles/10249647)

```
文章為自己的經驗與夥伴整理的內容，設計沒有標準答案，如有可以改進的地方，請告訴我，我會盡我所能的修改，謝謝大家～
```

---

大家好，今天要將 Istio 與 K8s 整合，整合之後會發現每個 K8s pod 都會從 1 個容器變為 2 個容器，

![](https://i.imgur.com/dXyh8GZ.png)

是因為容器都多了一個 envoy proxy 容器來管理流量，如圖，之後就可以再利用 envoy proxy 配合 Control Plane 來控制 Service Mesh。

## 安裝 Istio

[官網](https://istio.io/latest/docs/setup/getting-started/)已有相當清楚的教學，在我 Mac 上就是下載檔案並且加入執行的環境變數即可。

```bash
$ curl -L https://istio.io/downloadIstio | sh -
$ cd istio-1.7.3
$ export PATH=$PWD/bin:$PATH
```

## 開始整合

我們將以此[Example-Code](https://github.com/superj80820/2020-ithelp-contest)來說明，

Istio 至少需要 K8s 1.16 以上才支持，而之前的文章是使用 K8s 1.15，所以必須先刪除 minikube 的 K8s 1.15，

```bash
$ minikube delete
```

創建 1.16 的 K8s，並開起 ingress

```bash
$ minikube start --driver=hyperkit --kubernetes-version v1.16.0
$ minikube addons enable ingress
```

取得 minikube ip，

```bash
$ minikube ip
```

將特定`api.backend.com`與`web.backend.com`轉換成 minikube ip，以確保 ingress 正常運作，

![](https://i.imgur.com/mc9HuVS.png)

安裝 Istio 至 K8s，

```bash
$ istioctl install --set profile=demo
```

啟用`istio-injection`，這會使在建立 pod 的時候，自動在 pod 中加入 envoy proxy 容器，

```bash
$ kubectl label namespace default istio-injection=enabled
```

再透過 Helm 啟動之前，因為 K8s 1.16 與 1.15 有一些[breaking change](https://kubernetes.io/blog/2019/07/18/api-deprecations-in-1-16/)，所以必須要在 deployment config 中加入 selector，如圖，

![](https://i.imgur.com/IuxHnmo.png)

最後啟動 K8s 應用！

```yaml
helm install . --generate-name
```

![](https://i.imgur.com/QvobsJT.png)

Work!而紅匡處中容器多了 envoy proxy，但綠匡處會發現一直啟動不了 envoy proxy，這導致`api.backend.com`也沒辦法正常運作，

不過不用擔心，因為 Istio 本身就是設計每個容器都要有個 sidecar envoy proxy 了，實在沒必要再設計一個 envoy 的 service，我們直接使用 Istio 原生的功能來設計即可，將在之後介紹。

## 參考

- [Istio / Getting Started](https://istio.io/latest/docs/setup/getting-started/)
- [Istio / What is Istio?](https://istio.io/latest/docs/concepts/what-is-istio/)
- [Deprecated APIs Removed In 1.16: Here’s What You Need To Know | Kubernetes](https://kubernetes.io/blog/2019/07/18/api-deprecations-in-1-16/)
