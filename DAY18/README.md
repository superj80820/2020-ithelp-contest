本文章同時發佈於：

- [Github(包含程式碼)](https://github.com/superj80820/2020-ithelp-contest/blob/master/DAY18)
- [Medium](https://medium.com/%E9%AB%92%E6%A1%B6%E5%AD%90/day18-%E4%BA%86%E8%A7%A3-k8s-%E4%B8%AD%E7%9A%84-pod-service-deployment-92408f9244e1)
- [iT 邦幫忙](https://ithelp.ithome.com.tw/articles/10248019)

```
文章為自己的經驗與夥伴整理的內容，設計沒有標準答案，如有可以改進的地方，請告訴我，我會盡我所能的修改，謝謝大家～
```

---

![](https://i.imgur.com/sHZdiDx.png)

大家好，今天要介紹 K8s 到底有哪些 config yaml 檔，他們又是何種用途。

## K8s 的整體架構

![](https://i.imgur.com/rsBTtXL.png)

圖片來源: [Kubernetes, At A First Glance](http://nishadikirielle.blogspot.com/2016/02/kubernetes-at-first-glance.html)

K8s 有非常多的元件來建構豐富多樣的架構，但我們可已由高至低了解以下 4 個不同層級的重要元件，會比較好理解:

- Cluster: 管理多個 Master, Worker Node，可以理解為多個 VM 如何變為一個「大 VM」的方式
- Master Node: 實際的 VM，K8s 會由一個 Master Node 去管理底下多個 Worker Node
- Worker Node: 實際的 VM，裡頭管理著一個一個 Pod
- Pod: Pod 裡放著容器，容器通常為一個，但也可以多個

## Pod、Deployment、Service yaml config 介紹

Pod 可以理解為一個或多個 docker 容器怎麼啟動在 K8s 中，而在 kompose 轉換中沒有出現此 config 檔案，為什麼呢？

因為 Pod 之上有 Deployment，他除了描述 Pod 的行為外，還會描述`Pod在K8s上狀態發生變化時要對應的機制`。

而關於`如何連線存取這些Pod`，就要透過 Service 來設定

---

Service 與 Deployment 中有些 config 使用的欄位相同，分別是:

- apiVersion: 為描述此元件使用的 K8s API 的版本
- kind: 說明此元件屬於何種類型，有 Service、Deployment 等等
- metadata: 說明此元件的相關資訊
- metadata.labels: 可以將此元件分類在特定 label
- metadata.annotations: 說明不需要分類的相關資訊
- metadata.creationTimestamp: 此元件被創立的時間
- metadata.name: 此元件的名稱

Deployment 的欄位介紹:

```yaml
# server-service.yaml
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
  replicas: 1
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

- spec.replicas: 此 Pod 會在 K8s 有幾個橫向擴展(Horizontal Pod Autoscaler)，目前設定一個
- spec.strategy: 可以設定狀態變化對應機制的策略，例如 image 降版要要維持幾個 Pod 之類，這邊使用預設值
- spec.template.spec.containers: 設置容器
- spec.template.spec.containers.command: 容器的啟動 command
- spec.template.spec.containers.image: 容器使用的 image
- spec.template.spec.containers.ports: 容器使用的 port
- spec.template.spec.containers.restartPolicy: 容器是否無預期關閉後要重新啟動

Service 的欄位介紹:

```yaml
# server-service.yaml
apiVersion: v1
kind: Service
metadata:
  annotations:
    kompose.cmd: kompose convert
    kompose.version: 1.16.0 (0c01309)
  creationTimestamp: null
  labels:
    io.kompose.service: server
  name: server
spec:
  ports:
    - name: "6000"
      port: 6000
      targetPort: 6000
  selector:
    io.kompose.service: server
status:
  loadBalancer: {}
```

- spec.ports.ports: 說明對外可連入的 port 為何
- spec.ports.targetPort: 說明對外連入的 port 對應到 Pod 的哪個 port
- spec.ports.selector: 此規則要套用到哪個 label 上

## 參考

- [Kubernetes 基礎教學（一）原理介紹. Kubernetes 如何運作？什麼是 Pod？什麼是 Node？Master… | by 胡程維｜ Cheng-Wei Hu | Medium](https://medium.com/@C.W.Hu/kubernetes-basic-concept-tutorial-e033e3504ec0)
- [Kubernetes 基礎教學（二）實作範例：Pod、Service、Deployment、Ingress | by 胡程維｜ Cheng-Wei Hu | Medium](https://medium.com/@C.W.Hu/kubernetes-implement-ingress-deployment-tutorial-7431c5f96c3e)
- [Kubernetes 基礎教學（三）Helm 介紹與建立 Chart | by 胡程維｜ Cheng-Wei Hu | Medium](https://medium.com/@C.W.Hu/kubernetes-helm-chart-tutorial-fbdad62a8b61)
- [[Day 10] Kubernetes 世界不可缺少的 - Labels - iT 邦幫忙::一起幫忙解決難題，拯救 IT 人的一天](https://ithelp.ithome.com.tw/articles/10194613)
- [community/api-conventions.md at master · kubernetes/community](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md)
- [Kubernetes, at a first glance | Nishadi's Tech Blog](http://nishadikirielle.blogspot.com/2016/02/kubernetes-at-first-glance.html)
- [透過 Kubernetes Deployments 實現滾動升級 | Kubernetes](https://tachingchen.com/tw/blog/kubernetes-rolling-update-with-deployment/)
