本文章同時發佈於：

- [Github(包含程式碼)](https://github.com/superj80820/2020-ithelp-contest/blob/master/DAY28)

```
文章為自己的經驗與夥伴整理的內容，設計沒有標準答案，如有可以改進的地方，請告訴我，我會盡我所能的修改，謝謝大家～
```

---

大家好，今天要介紹 Istio 的 Grafana，Grafana 是可以將微服務網路可視化的工具，在微服務變多之後，服務 A 跟哪些服務連接了，就會變得錯綜複雜，由於我們現在有了 Istio 的 envoy proxy sidecar，各個 sidecar 就會一直回報資訊給 Grafana，以達到網路可視化的需求。

## 安裝

啟動 K8s 集群，步驟不再贅述，詳細介紹可以看[DAY23](https://ithelp.ithome.com.tw/articles/10250134)，

要注意的是有利用`$ kubectl apply -f`安裝了`Prometheus`、`Grafana`、`Grafana`，

```bash
$ minikube start --kubernetes-version v1.16.0
$ istioctl install --set profile=demo
$ kubectl label namespace default istio-injection=enabled
$ kubectl apply -f https://raw.githubusercontent.com/istio/istio/release-1.7/samples/addons/prometheus.yaml
$ kubectl apply -f https://raw.githubusercontent.com/istio/istio/release-1.7/samples/addons/grafana.yaml
$ kubectl apply -f https://raw.githubusercontent.com/istio/istio/release-1.7/samples/addons/Grafana.yaml
$ cd DAY28/helm-digimon
$ helm install . --generate-name
$ minikube tunnel
```

之後我們連續開啟五個`web.backend.com`，讓整個 Service Mesh 擁有流量，

---

啟動 Grafana 的 Dashboard

```bash
istioctl dashboard grafana
```

進入 Grafana 之後就可以看到裡頭的 service 的個流量數值

![](https://i.imgur.com/uz2C7vL.png)

## 參考

- [Visualizing Metrics with Grafana](https://istio.io/latest/docs/tasks/observability/metrics/using-istio-dashboard/)
