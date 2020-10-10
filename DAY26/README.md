本文章同時發佈於：

- [Github(包含程式碼)](https://github.com/superj80820/2020-ithelp-contest/blob/master/DAY26)

```
文章為自己的經驗與夥伴整理的內容，設計沒有標準答案，如有可以改進的地方，請告訴我，我會盡我所能的修改，謝謝大家～
```

---

大家好，接下來要介紹 Istio DestinationRule 元件實作的部分，

## 實作

以下的實作範例都在[Code-Example](https://github.com/superj80820/2020-ithelp-contest)中，

code 中以 server 的 pods 改為兩個，

![](https://i.imgur.com/EleGXnO.png)

並且新增了一個 DestinationRule 元件在`DAY26/helm-digimon/templates/load-balancer.yaml`

```yaml
apiVersion: networking.istio.io/v1alpha3
kind: DestinationRule
metadata:
  name: server
spec:
  host: server
  trafficPolicy:
    loadBalancer:
      simple: ROUND_ROBIN
```

- `host`: 會對 server 這個 service 來做流量處理
- `trafficPolicy.loadBalancer.simple`: 選擇流量處理的策略，其他不同的策略[在此](https://istio.io/latest/docs/reference/config/networking/destination-rule/#LoadBalancerSettings-SimpleLB)

啟動 K8s 集群，步驟不再贅述，詳細介紹可以看[DAY23](https://ithelp.ithome.com.tw/articles/10250134)

```bash
$ minikube start --kubernetes-version v1.16.0
$ istioctl install --set profile=demo
$ kubectl label namespace default istio-injection=enabled
$ cd DAY26/helm-digimon
$ helm install . --generate-name
$ minikube tunnel
```

再來我們要觀看 pods 是不是有正確分流，使用`$ kubectl get pods`拿取全部`pods`，

![](https://i.imgur.com/K6DelGm.png)

使用`$ kubectl logs`來取得 pods 的 log，以我來說我會是以下指令:

```bash
$ kubectl logs -f server-66db48686-f4xm5
$ kubectl logs -f server-66db48686-sdpfs
```

![](https://i.imgur.com/FXI0PY9.gif)

兩個 pods 已經成功被分流。

---

如果想要看看沒有 DestinationRule 元件會產生什麼狀況，可以把`DAY26/helm-digimon/templates/load-balancer.yaml`刪除，再重新以上步驟創建一個 K8s 集群

![](https://i.imgur.com/xpC33lG.gif)

我們會看到流量只有導到左邊的 pods，造成左邊很忙右邊很閒的狀況 XD。

## 參考

- [Istio / Destination Rule](https://istio.io/latest/docs/reference/config/networking/destination-rule/)
- [Istio 基礎 — gRPC 負載均衡. 什麼是 Istio | by Alan Chen | getamis | Medium](https://medium.com/getamis/istio-%E5%9F%BA%E7%A4%8E-grpc-%E8%B2%A0%E8%BC%89%E5%9D%87%E8%A1%A1-d4be0d49ee07)
- [Using Istio to load-balance internal gRPC services  |  Solutions](https://cloud.google.com/solutions/using-istio-for-internal-load-balancing-of-grpc-services)
