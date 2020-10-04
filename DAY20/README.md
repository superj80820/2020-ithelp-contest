本文章同時發佈於：

- [Github(包含程式碼)](https://github.com/superj80820/2020-ithelp-contest/blob/master/DAY20)
- [iT 邦幫忙](https://ithelp.ithome.com.tw/articles/10249122)

```
文章為自己的經驗與夥伴整理的內容，設計沒有標準答案，如有可以改進的地方，請告訴我，我會盡我所能的修改，謝謝大家～
```

---

大家好，今天我們要介紹 Helm，

昨天已經將 K8s 的 Pod、Service、Deployment、Ingress 元件都介紹完畢，容器的管理都變得更加容易，但各個元件的 config yaml 檔將會變得如下圖這麼多，

![](https://i.imgur.com/bCoM8Pt.png)

這麼多的元件導致每次 K8s 啟動時都要下以下指令：

```bash
$ kubectl create -f db-deployment.yaml,db-service.yaml,envoy-deployment.yaml,envoy-service.yaml,server-deployment.yaml,server-service.yaml,weather-deployment.yaml,weather-service.yaml,web-deployment.yaml,web-service.yaml
$ kubectl create -f ingress.yaml
```

就管理上還是麻煩了些，而 Helm 就是將這些 config yaml 檔在包裝起來的工具，讓後端部署更加的`一鍵解決`XD。

## Helm 安裝

Mac 安裝可以使用 Homebrew 安裝，

```bash
brew install helm
```

而其他的平台也相當簡單，可以參考[官網](https://github.com/helm/helm#install)

安裝完後，記得將 Helm repostory URL 添加至 Helm，這就像 Mac 的 Homebrew、Ubuntu 的 apt 的 library URL

```bash
helm repo add stable https://kubernetes-charts.storage.googleapis.com/
```

## 先 run 起來

Clone [Example-Code](https://github.com/superj80820/2020-ithelp-contest)，進入到 Helm 的資料夾，

```bash
$ cd DAY20/helm-digimon
```

將 Helm 啟動，

```bash
$ helm install . --generate-name
```

此時可以下以下指令，會發現 Helm 將我們前幾篇文章的元件一次建立完畢，

```bash
$ kubectl get all
```

![](https://i.imgur.com/NAOEjvm.png)

而開啟`web.digimon.com`，work！

![](https://i.imgur.com/fgkGYAx.png)

## 實作

建立一 helm chart，

```bash
$ helm create helm-digimon
```

指令下完後，會建立以下資料夾，

![](https://i.imgur.com/X5wvfSt.png)

- charts: 放置其他 chart 的資料夾，由於這次的實作較簡單，所以並沒有用到其他 chart
- templates: 放置 K8s 的元件
- .helmignore: 要忽略的檔案
- Chart.yaml: 說明此 chart 的 metadata，有名稱與版本等等
- values.yaml: 可將 K8s 元件的參數抽象出來，並在 values.yaml 統一控管

由於 templates 中的元件我都將它刪除了，放入前幾篇文章的元件，會如下，

![](https://i.imgur.com/gESTuC0.png)

並把 deployment 裡頭的`replicaCount`抽象出來，統一由 value.yaml 來控管，以 server-deployment.yaml 來說會如下：

```yaml
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  annotations:
    kompose.cmd: kompose convert
    kompose.version: 1.16.0 (0c01309)
  creationTimestamp: null
  labels:
    io.kompose.service: server
  name: server
spec:
  replicas: { { .Values.replicaCount } }
  strategy: {}
  template:
    metadata:
      creationTimestamp: null
      labels:
        io.kompose.service: server
    spec:
      containers:
        - command:
            - go
            - run
            - cmd/main.go
          image: superj80820/digimon-service
          name: server
          ports:
            - containerPort: 6000
          resources: {}
      restartPolicy: Always
status: {}
```

## 參考

- [Kubernetes 基礎教學（二）實作範例：Pod、Service、Deployment、Ingress | by 胡程維｜ Cheng-Wei Hu | Medium](https://medium.com/@C.W.Hu/kubernetes-implement-ingress-deployment-tutorial-7431c5f96c3e)
- [Kubernetes 基礎教學（三）Helm 介紹與建立 Chart | by 胡程維｜ Cheng-Wei Hu | Medium](https://medium.com/@C.W.Hu/kubernetes-helm-chart-tutorial-fbdad62a8b61)
