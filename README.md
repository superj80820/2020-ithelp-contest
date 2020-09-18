# 2020-ithelp-contest

## 被選召的 Gopher 們，從零開始探索 Golang, Istio, K8s 數碼微服務世界

從單體服務一路到現在的微服務，Backend 經歷了不可思議的變遷，而 Golang 這短小精悍的語言就好像剛進到數碼寶貝世界的孩子，雖然一個人做的事情不多，但是只要相信 Kubernetes、Istio 等等數碼獸，就可以改變世界，本系列會以初學的角度出發，希望我與大家能有機會一起獲得微服務數碼徽章。

## 目錄

1. [DAY1 - 在開始數碼微服務之旅前](https://github.com/superj80820/2020-ithelp-contest/blob/master/DAY01)
2. [DAY2 - 微服務真的有那麼好？你的 Backend 有需要進化成他嗎？](https://github.com/superj80820/2020-ithelp-contest/blob/master/DAY02)
3. [DAY3 - Docker 在手，服務帶著走 - Docker 篇](https://github.com/superj80820/2020-ithelp-contest/blob/master/DAY03)
4. [DAY4 - 今晚，我想來點 Golang Server 加 PostgreSQL - Docker-Compose 篇](https://github.com/superj80820/2020-ithelp-contest/blob/master/DAY04)

## 想法

- 將數碼蛋傳送到第二個微服務: 神聖計畫吧！ - gRPC 篇 part 1(單向)
- 將數碼蛋傳送到第二個微服務: 神聖計畫吧！ - gRPC 篇 part 2(雙向)
- 將數碼蛋傳送到第二個微服務: 神聖計畫吧！ - gRPC 篇 part 3(中間件)
- 將 gRPC 帶回人類的前端世界 - gRPC-Web 篇
- 連接不同世界的橋樑: envoy
- Install k8s
- Why need Istio
- Install Istio
- 部署 orService discover?
- Istio ziping
- nats
- redis
- scaling
- k8s auto update

## 參考文章

- [Istio Observability with Go, gRPC, and Protocol Buffers-based Microservices | Programmatic Ponderings](https://programmaticponderings.com/2019/04/17/istio-observability-with-go-grpc-and-protocol-buffers-based-microservices/)
- [k8s-istio-observe-backend/main.go at master · garystafford/k8s-istio-observe-backend](https://github.com/garystafford/k8s-istio-observe-backend/blob/master/services/service-b/main.go)
- [Kubernetes 基礎教學（二）實作範例：Pod、Service、Deployment、Ingress | by 胡程維｜ Cheng-Wei Hu | Medium](https://medium.com/@C.W.Hu/kubernetes-implement-ingress-deployment-tutorial-7431c5f96c3e)
- [Kubernetes 應用 Istio 打造精準導流與更版. Istio、Kubernetes、Gateway、VirtualService、… | by Tass | Medium](https://medium.com/@tasslin/kubernetes%E6%87%89%E7%94%A8istio%E6%89%93%E9%80%A0%E7%B2%BE%E6%BA%96%E5%B0%8E%E6%B5%81%E8%88%87%E6%9B%B4%E7%89%88-9206f7d326ff)
- [Tass (@tasslin) Articles – Medium](https://medium.com/@tasslin)
- [Istio step-by-step Part 01 — Introduction to Istio | by Nethmini Romina | FAUN | Medium](https://medium.com/faun/istio-step-by-step-part-01-introduction-to-istio-b9fd0df30a9e)
- [itnext.io](https://itnext.io/micro-in-action-getting-started-a79916ae3cac)
- [Micro In Action, Part 2: An Ultimate Guide For Bootstrap | by Che Dan | ITNEXT](https://itnext.io/micro-in-action-part-2-71230f01d6fb)
- [Micro In Action, Part 3: Calling a Service | by Che Dan | ITNEXT](https://itnext.io/micro-in-action-part-3-calling-a-service-55d865928f11)
- [Micro In Action, Part4：Pub/Sub. This is the 4th article in the series… | by Che Dan | ITNEXT](https://itnext.io/micro-in-action-part4-pub-sub-564f3b054ecd)
- [go-kit/kit: A standard library for microservices.](https://github.com/go-kit/kit)
- [Mastering Wire. Introducing the concept, basic usage… | by Che Dan | ITNEXT](https://itnext.io/mastering-wire-f1226717bbac)
- [micro/go-micro: A Go standard library for microservices](https://github.com/micro/go-micro)
- [Microservices with go-kit. Part 1 - DEV Community 👩‍💻👨‍💻](https://dev.to/plutov/microservices-with-go-kit-part-1-13dd)
- [Fix ending separator for the gopath by kujtimiihoxha · Pull Request #31 · kujtimiihoxha/kit](https://github.com/kujtimiihoxha/kit/pull/31)
- [Build go kit Microservices at Kubernetes with ease - Kai-Chu Chung - HackMD](https://hackmd.io/@gdg-taipei/BJ1ImG03B)
- [Istio / 架构](https://istio.io/latest/zh/docs/ops/deployment/architecture/)
- [go-kit/kit: A standard library for microservices.](https://github.com/go-kit/kit)
- [Go Microservices with Go kit: Introduction | by Shiju Varghese | Medium](https://medium.com/@shijuvar/go-microservices-with-go-kit-introduction-43a757398183)
- [shijuvar/gokit-examples: Examples for building microservices with Go kit (gokit.io)](https://github.com/shijuvar/gokit-examples)
- [如何使用 Go kit 工具包编写微服务 - Go 语言中文网 - Golang 中文社区](https://studygolang.com/articles/21378)
- [How to write a microservice in Go with Go kit - DEV Community 👩‍💻👨‍💻](https://dev.to/napolux/how-to-write-a-microservice-in-go-with-go-kit-a66)
- [如何在 Go with Go 套件中編寫微服務–編碼](https://coding.napolux.com/how-to-write-a-microservice-in-go-with-go-kit/)
- [napolux/go-kit-microservice-example-tutorial-99999: A little "Restful" Go kit microservice I built to learn a bit more about Go kit](https://github.com/napolux/go-kit-microservice-example-tutorial-99999)
- [Go kit-示例](https://gokit.io/examples/)
- [gokit.io](https://gokit.io/examples/stringsvc.html#calling-other-services)
- [kit/instrumenting.go at master · go-kit/kit](https://github.com/go-kit/kit/blob/master/examples/shipping/tracking/instrumenting.go)
- [Microservices with go-kit. Part 2 - DEV Community 👩‍💻👨‍💻](https://dev.to/plutov/packagemain-13-microservices-with-go-kit-part-2-4lgh)
- [packagemain/13-go-kit-2 at master · plutov/packagemain](https://github.com/plutov/packagemain/tree/master/13-go-kit-2)
- [帶有 Go-kit 的微服務。 第 1 部分-DEV 社區 👩‍💻👨‍💻](https://dev.to/plutov/microservices-with-go-kit-part-1-13dd)
- [Go kit - Frequently asked questions](https://gokit.io/faq/#what-is-go-kit)
- [etcd-io/etcd: Distributed reliable key-value store for the most critical data of a distributed system](https://github.com/etcd-io/etcd)
- [kit/transport.go at master · go-kit/kit](https://github.com/go-kit/kit/blob/master/examples/shipping/booking/transport.go)
- [micro/go-micro: A Go standard library for microservices](https://github.com/micro/go-micro)
- [itnext.io](https://itnext.io/micro-in-action-part-3-calling-a-service-55d865928f11)
- [itnext.io](https://itnext.io/micro-in-action-part6-service-discovery-f988988e5936)
- [zhuanlan.zhihu.com](https://zhuanlan.zhihu.com/p/79897640)
