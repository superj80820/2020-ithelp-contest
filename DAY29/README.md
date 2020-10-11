本文章同時發佈於：

- [Github(包含程式碼)](https://github.com/superj80820/2020-ithelp-contest/blob/master/DAY29)

```
文章為自己的經驗與夥伴整理的內容，設計沒有標準答案，如有可以改進的地方，請告訴我，我會盡我所能的修改，謝謝大家～
```

---

大家好，今天要介紹 Istio 的 Jaeger，Jaeger 可以對請求流量進行跟蹤，是微服務中複雜請求的追蹤方案

## 安裝

啟動 K8s 集群，步驟不再贅述，詳細介紹可以看[DAY23](https://ithelp.ithome.com.tw/articles/10250134)，

要注意的是有利用`$ kubectl apply -f`安裝了`Prometheus`、`Grafana`、`Grafana`、`Jaeger`，

```bash
$ minikube start --kubernetes-version v1.16.0
$ istioctl install --set profile=demo
$ kubectl label namespace default istio-injection=enabled
$ kubectl apply -f https://raw.githubusercontent.com/istio/istio/release-1.7/samples/addons/jaeger.yaml
$ kubectl apply -f https://raw.githubusercontent.com/istio/istio/release-1.7/samples/addons/prometheus.yaml
$ kubectl apply -f https://raw.githubusercontent.com/istio/istio/release-1.7/samples/addons/grafana.yaml
$ kubectl apply -f https://raw.githubusercontent.com/istio/istio/release-1.7/samples/addons/Grafana.yaml
$ cd DAY29/helm-digimon
$ helm install . --generate-name
$ minikube tunnel
```

在 Istio 預設值`values.pilot.traceSampling`是設為 1%，意思是 100 個 requests 採樣一次，所以我們必須打開網頁並將 `Create` request 複製起來，利用 bash 打 100 次，進入`web.backend.com`之後開啟`開發者模式`，

![](https://i.imgur.com/EeGRhyf.jpg)

並以 for 迴圈請求 100 次

```bash
$ for i in $(seq 1 100); do {複製的cURL} done
```

以我來說會如下:

```bash
$ for i in $(seq 1 100); do curl 'http://api.backend.com/digimon.Digimon/Create' \
  -H 'Connection: keep-alive' \
  -H 'Accept: application/grpc-web-text' \
  -H 'X-User-Agent: grpc-web-javascript/0.1' \
  -H 'X-Grpc-Web: 1' \
  -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36' \
  -H 'Content-Type: application/grpc-web-text' \
  -H 'Origin: http://web.backend.com' \
  -H 'Referer: http://web.backend.com/' \
  -H 'Accept-Language: en-US,en;q=0.9,zh-TW;q=0.8,zh;q=0.7' \
  --data-binary 'AAAAAAgKBkFndW1vbg==' \
  --compressed \
  --insecure; done
```

---

開啟 Jaeger

```bash
$ istioctl dashboard jaeger
```

選取`istio-ingressgateway`之後點選搜尋，會發現總共有 100 個追蹤，

![](https://i.imgur.com/BUE6uUt.png)

點選其中一個追蹤，會列出此請求在經過了哪些 service，因為 `Create` request 本身沒有經過不同的 service 所以 dashboard 只顯示了一個 service，點選 service 就會在顯示此 request 在此 service 上的 tag 狀態，

![](https://i.imgur.com/LoJ2UjL.png)

以官方的例子，如果是眾多 services，request 在不同 service 走動花費的時間就會如下顯示，

![](https://i.imgur.com/goBCDZH.png)

## 參考

- [Istio / Jaeger](https://istio.io/latest/docs/ops/integrations/jaeger/#installation)
- [Istio / Jaeger](https://istio.io/latest/docs/tasks/observability/distributed-tracing/jaeger/)
- [jaegertracing/jaeger: CNCF Jaeger, a Distributed Tracing Platform](https://github.com/jaegertracing/jaeger)
- [Day 26 - 安裝使用分散式跟蹤系統 - Jaeger - iT 邦幫忙::一起幫忙解決難題，拯救 IT 人的一天](https://ithelp.ithome.com.tw/articles/10207800)
