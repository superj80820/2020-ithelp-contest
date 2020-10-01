本文章同時發佈於：

- [Github(包含程式碼)](https://github.com/superj80820/2020-ithelp-contest/blob/master/DAY17)
- [Medium]()
- [iT 邦幫忙](https://ithelp.ithome.com.tw/articles/10247631)

```
文章為自己的經驗與夥伴整理的內容，設計沒有標準答案，如有可以改進的地方，請告訴我，我會盡我所能的修改，謝謝大家～
```

---

![](https://i.imgur.com/YDTq5K0.png)

大家好，這次要介紹 kompose，我們將先 run 起 K8s，再來介紹 K8s 的細節，這樣可能比較好讓大家理解(包括我 XD)

## 安裝相關套件

- [minikube](https://kubernetes.io/docs/tasks/tools/install-minikube/): 可以在本機電腦安裝一個簡單的 K8s 環境
- [kompose](https://kubernetes.io/zh/docs/tasks/configure-pod-container/translate-compose-kubernetes/): 將 docker-compose 檔無痛轉移至 K8s 的 yaml 檔

![](https://i.imgur.com/4HZwrV1.png)

圖片來源: [deploy-kubernetes-the-ultimate-guide](https://platform9.com/docs/deploy-kubernetes-the-ultimate-guide/)

minikube 在本機電腦會安裝一個 VM，以模擬叢集(Cluster)的環境，以減少 K8s 的配置時間成本。

## 實作

clone 範例 [Github-Example-Code](https://github.com/superj80820/2020-ithelp-contest)，並進入此篇文章的範例 folder

```bash
$ cd DAY17
```

啟動 minikube `1.15.0` 的 K8s 版本，這個版本與 kompose 相容性較好

```bash
$ minikube start --kubernetes-version v1.15.0
```

使用 kompose 來啟動 K8s 的相關 config

```bash
$ kompose up
```

瀏覽實際的網頁

```bash
$ minikube service web
```

此時會看到 web 的相關 IP

![](https://i.imgur.com/T62NtcY.png)

打開瀏覽器，Work!

![](https://i.imgur.com/ferJqWn.png)

## 轉換出 yaml 檔

![](https://i.imgur.com/FTdisWd.png)

![](https://i.imgur.com/sHZdiDx.png)

這些就是執行`kompose up`實際的 config yaml 檔，而他們代表的意思將在之後文章介紹，謝謝你的閱讀～

## 參考

- [Dockerfile 中的 ENTRYPOINT. ENTRYPOINT 是 Dockerfile 定義的一個指令，他的作用類似於… | by Justin Chien | Medium](https://medium.com/@xyz030206/dockerfile-%E4%B8%AD%E7%9A%84-entrypoint-9653c3b2d2f8)
- [将 Docker Compose 文件转换为 Kubernetes 资源 | Kubernetes](https://kubernetes.io/zh/docs/tasks/configure-pod-container/translate-compose-kubernetes/)
- [Installing Kubernetes with Minikube | Kubernetes](https://kubernetes.io/docs/setup/learning-environment/minikube/)
- [Install Minikube | Kubernetes](https://kubernetes.io/docs/tasks/tools/install-minikube/)
- [Deploy Kubernetes: The Ultimate Guide - Platform9](https://platform9.com/docs/deploy-kubernetes-the-ultimate-guide/)
