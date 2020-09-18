本文章同時發佈於：

- [Github(包含程式碼)]()
- [Medium]()
- [iT 邦幫忙]()

```
文章為自己的經驗與夥伴整理的內容，設計沒有標準答案，如有可以改進的地方，請告訴我，我會盡我所能的修改，謝謝大家～
```

---

[//]: # "york TODO: 新增 DAY06 連結"

大家好，繼昨天[DAY06]()的介紹後，相信大家已經對 Clean Architecture 的稍有概念了，接下來將介紹實作的部分，相信會讓各位更有感覺。

## 安裝 Golang

關於安裝官網寫得很清楚，在此貼上[連結](https://golang.org/doc/install)就不多做介紹了。

## 透過[swagger-generator](https://github.com/swagger-api/swagger-codegen)來產生 Server 介面

[//]: # "york TODO: 補充連結"

進到[DAY07-example]()的資料夾，使用 docker 運行以下指令:

```bash
$ docker run --rm -v ${PWD}:/local swaggerapi/swagger-codegen-cli-v3 generate \
    -i /local/swagger.yaml \
    -l go-server \
    -o /local/go-server
```

swagger-generator 會自動產生以下的 code:

![](https://i.imgur.com/0Vy5hSS.png)

進到`go-server`資料夾底下創建`go module`檔，並把專案名稱設為`go-server`，比較常見的做法是取名為 git 的網址，不過為了簡單介紹我們直接取`go-server`。`go module`是 Golang 用來管理套件的方式，類似於 Node.js 的 package.json，

```bash
$ cd go-server
$ go mod init go-server
```

把`main.go` import 的路徑改為 `go module` 的路徑，才能讓 go 檔案都吃得到，

![](https://i.imgur.com/0gQvazP.png)

運行一下試看看 Server 有沒有跑起來吧～

```bash
$ go run main.go
```

![](https://i.imgur.com/7RCUTAD.png)

Work!

## 參考

- [bxcodec/go-clean-arch](https://github.com/bxcodec/go-clean-arch)
- [Trying Clean Architecture on Golang](https://medium.com/hackernoon/golang-clean-archithecture-efd6d7c43047)
