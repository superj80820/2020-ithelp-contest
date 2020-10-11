本文章同時發佈於：

- [Github(包含程式碼)](https://github.com/superj80820/2020-ithelp-contest/blob/master/DAY30)

```
文章為自己的經驗與夥伴整理的內容，設計沒有標準答案，如有可以改進的地方，請告訴我，我會盡我所能的修改，謝謝大家～
```

---

大家好，今天要介紹 Istio 的 Prometheus，Prometheus 是一個時序資料庫，可以對於流量請求可以設定不同的 Alert 與各種圖形監控

## 安裝

啟動 K8s 集群，步驟不再贅述，詳細介紹可以看[DAY23](https://ithelp.ithome.com.tw/articles/10250134)，

要注意的是有利用`$ kubectl apply -f`安裝了`Prometheus`、`Grafana`、`Grafana`、`prometheus`，

```bash
$ minikube start --kubernetes-version v1.16.0
$ istioctl install --set profile=demo
$ kubectl label namespace default istio-injection=enabled
$ kubectl apply -f https://raw.githubusercontent.com/istio/istio/release-1.7/samples/addons/jaeger.yaml
$ kubectl apply -f https://raw.githubusercontent.com/istio/istio/release-1.7/samples/addons/prometheus.yaml
$ kubectl apply -f https://raw.githubusercontent.com/istio/istio/release-1.7/samples/addons/grafana.yaml
$ kubectl apply -f https://raw.githubusercontent.com/istio/istio/release-1.7/samples/addons/Grafana.yaml
$ cd DAY30/helm-digimon
$ helm install . --generate-name
$ minikube tunnel
```

我們將網頁並將 `Create` request 複製起來，利用 bash 打 100 次，進入`web.backend.com`之後開啟`開發者模式`，

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

開啟 Prometheus

```bash
$ istioctl dashboard prometheus
```

輸入`istio_requests_total`再點選 Execute，會看到所有的流量請求

![](https://i.imgur.com/HVRVEWO.png)

而點選左邊的 Graph 會看到流量的線條圖

![](https://i.imgur.com/QefaiXc.png)

---

終於 30 天完賽了，在後面這幾天監控會發現，其實安裝與開啟部分都差不多，這是 Istio 很大的好處，即`各種監控已經整合至平台層，不需再另外設計`，沒有特別寫完賽文是因為，最後這幾天都是初探這些工具，在未來我希望可以在後續文章補齊使用的心得與感想，謝謝各位一路陪伴～

## 參考
