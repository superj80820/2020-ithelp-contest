本文章同時發佈於：

- [Github(包含程式碼)](https://github.com/superj80820/2020-ithelp-contest/blob/master/DAY16)
- [Medium](https://medium.com/%E9%AB%92%E6%A1%B6%E5%AD%90/day16-kubernetes-%E7%AE%A1%E7%90%86%E4%B8%80%E5%88%87%E5%BE%AE%E6%9C%8D%E5%8B%99%E7%9A%84%E4%B8%96%E7%95%8C%E6%A8%B9-f46af8e1f424)
- [iT 邦幫忙](https://ithelp.ithome.com.tw/articles/10246632)

```
文章為自己的經驗與夥伴整理的內容，設計沒有標準答案，如有可以改進的地方，請告訴我，我會盡我所能的修改，謝謝大家～
```

---

![](https://i.imgur.com/BPMtKcf.png)

大家好，在前面 15 天講解完`應用層`相關的微服務後，現在終於要來介紹`平台層`！

而[Kubernetes](https://kubernetes.io/)(K8s)即是這次主角，他到底跟 Docker 是什麼樣的關係？你的後端需要他嗎？

我會以一個工程師，從`純Docker後端慢慢到發現需要K8s哪些好用的功能`，來說明 K8s 解決的問題。

## code 即是後端的架構

> 當後端容器越來越多，就會需要可以管理大量容器的方案

一開始使用 Docker 時，後端的容器不多，可以單純使用 docker-compose 即可解決多個容器互動的需求，但在容器越來越多時，會產生以下問題:

- 容器的連接變得複雜
- 容器的更新須寫腳本維護
- 容器是否掛了難以監控
- 容器如果流量需求增高，要自動增加容器來緩解流量
- 容器降版/升版需要腳本來控制

關於這些問題，我曾經寫過腳本來維護，但當服務越來越大，腳本的需求就越來越多，需要非常多成本。而 K8s 提供了像以下`config`的 yaml 形式來控制以上問題，讓一切架構以`config`來設計並非自己造輪子寫腳本來維護

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: site
  labels:
    app: web
spec:
  containers:
    - name: front-end
      image: nginx
      ports:
        - containerPort: 80
```

## K8s 架構

![](https://i.imgur.com/H3HbcvY.png)

圖片來源: [Project Management & Technology Fusion: Using Kubernetes (K8S) to manage Container cluster (Docker)](http://pmtechfusion.blogspot.com/2017/07/using-kubernetes-k8s-to-manage.html)

K8s 會先啟動一個 Master node，底下會有許多的 Worker Node，而 Worker node 裡頭會有許多的容器，K8s 主要的目標就是管理這些容器。

K8s 主要就是把多台機器加以抽象，使他們就像一台大機器工你操作。

而 K8s 最大目標是管理容器，而不管理運行容器的 VM，不過他依然有提供相關方案來管理 VM，但你使用的平台必須要有支援`Cluster Autoscaler(CA)`

## HPA、VPA、CA

K8s 有提供三種方法來管理`容器對流量的需求`

- Horizontal Pod Autoscaler(HPA): 以增加容器的方式來提供流量需求
- Vertical Pod Autoscaler(VPA): 以分配容器所需的 CPU/Memory 來提供流量需求
- Cluster Autoscaler(CA): 已增加整個叢集來提供流量需求

## 參考

- [Kubernetes Vs Docker | Sumo Logic](https://www.sumologic.com/blog/kubernetes-vs-docker/)
- [Why (and when) you should use Kubernetes | Hacker Noon](https://hackernoon.com/why-and-when-you-should-use-kubernetes-8b50915d97d8)
- [Do you really need Kubernetes?](https://blog.thundra.io/do-you-really-need-kubernetes)
- [Why is Kubernetes getting so popular? - Stack Overflow Blog](https://stackoverflow.blog/2020/05/29/why-kubernetes-getting-so-popular/)
- [Create Horizontally Auto-Scaling Cluster on AWS | Platform9 Documentation](https://docs.platform9.com/kubernetes/create-cluster-aws-horizontal-autoscale/)
- [你必知的 Kubernetes 自动缩放 · Service Mesh|服务网格中文社区](https://www.servicemesher.com/blog/k8s-autoscaling-all-you-need-to-know/)
- [Kuberenetes Autoscaling 相關知識小整理 • Weihang Lo](https://weihanglo.tw/posts/2020/k8s-autoscaling/)
- [Project Management & Technology Fusion: Using Kubernetes (K8S) to manage Container cluster (Docker)](http://pmtechfusion.blogspot.com/2017/07/using-kubernetes-k8s-to-manage.html)
