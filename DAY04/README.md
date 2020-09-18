本文章同時發佈於：

- [Github(包含程式碼)](https://github.com/superj80820/2020-ithelp-contest/blob/master/DAY04)
- [Medium]()
- [iT 邦幫忙]()

```
文章為自己的經驗與夥伴整理的內容，設計沒有標準答案，如有可以改進的地方，請告訴我，我會盡我所能的修改，謝謝大家～
```

---

大家好，繼昨天[DAY03](https://github.com/superj80820/2020-ithelp-contest/blob/master/DAY03)Docker 的介紹後，接下來要介紹如何 Docker-Compose 來管理多個環境。

## 為何需要 Docker-Compose ?

![](./digimon-service.drawio.png)

還記得昨天的 Service 架構圖嗎？裡頭除了 PostgreSQL 以外還有 Golang Server，我們一樣可以透過 Docker 來對 Golang Server 的環境做封裝，以提供所有的開發人員都在同一個環境開發。

而當有一堆 Docker 時，我們希望有個方法可以把

> Docker 的啟動方法、環境變數、容器間的關係寫成成配置文件，並利用此文件簡單的啟動所有容器

這時 Docker-Compose 就出馬了！

## 統一 Golang 環境的好處

可以避免不同人員裝了不同版本的 Golang，導致程式運行狀況有所差異的問題。不同開發人員也不用為了要安裝相同的 Golang 而飽受`更新`、`降版`地獄。

你可能會問，什麼時候會需要這麼多版本的 Golang？

答案是當管理一堆專案時。

以前我管理一堆 Python 專案時，Python2 跟 Python3 的歷史淵源導致套件版本相依問題非常嚴重，導致我與夥伴們是`三天一小亂，五天直接重灌`，而後來改用 Docker 之後這個問題就完美的解決了。

雖然 Golang 目前版本差異不大應該不太會有大問題，但未來怎麼樣誰知道呢？你還是可以試看看這個不錯的選擇。

## 安裝 Docker-Compose

在官網上都有把安裝的[教學](https://docs.docker.com/compose/install/)寫得很清楚了，這篇文章就關注實作而不是安裝了。

## Docker-Compose 檔案配置([Example code](https://github.com/superj80820/2020-ithelp-contest/blob/master/DAY04))

製作 Golang Server 的 `Dockerfile`，

```Dockerfile
# Golang Server's Dockerfile
FROM golang:1.14.6-alpine

RUN apk add --no-cache git

RUN go get github.com/lib/pq
```

- `FROM golang:1.14.6-alpine`: 選用 Golang 環境的版本
- `RUN apk add --no-cache git`: 因為 Golang 下載模組需要 Git，我們透過 apk 來安裝
- `RUN go get github.com/lib/pq`: 安裝 Golang 連線 PostgreSQL 的模組

製作`docker-compose.yaml`檔，

```docker-compose.yaml
version: "3.5"

services:
  server:
    build:
      context: ./docker/golang
      dockerfile: Dockerfile
    working_dir: /server
    volumes:
      - .:/server
    ports:
      - "5000:5000"
    depends_on:
      - db
    entrypoint: go run main.go
    restart: always
  db:
    image: postgres:12.4-alpine
    ports:
      - "5432:5432"
    environment:
      - POSTGRES_USER=user
      - POSTGRES_PASSWORD=mysecretpassword
      - PGDATA=/var/lib/postgresql/data/pgdata
    restart: always
```

- `version: "3.5"`: 選定 docker-compose 的版本，每個版本提供的 API 方法有所差異。
- `services`: 此欄位底下會有所有的容器，以下分別有`server`與`db`兩個 容器。
- `build`: 說明此容器要使用特定 Dockerfile 來 build，`context`為檔案目錄，`dockerfile`為 Dockerfile 的名字。
- `working_dir`: 指定 docker 啟動時所在的目錄，如果目錄不存在會自動創建一個。
- `volumes`: 將本機檔案掛載至 docker 內部，本機檔案更新時 docker 內部的掛載檔案也會更新。
- `ports`: 將本機的 port 做 mapping 與 docker 內部的 poart。
- `depends_on`: 說明 a 容器與 b 容器有相關，會等到 b 容器啟動完畢後，再啟動 a 容器。
- `entrypoint`: 指定 docker 啟動時的預設指令。
- `restart`: 當容器不正常關閉時，會重新啟動容器。
- `image`: 如果不使用 Dockerfile 來建立容器，你可以直接使用 docker image 來啟動容器。
- `environment`: 指定容器內的環境變數。

可以看到`docker-compose.yaml`很方便，把昨天文章中要輸入的`environment`變數都變成了設定檔。

透過`working_dir`配合`volumes`我們可以把當下目錄(即是`.`目錄)的檔案掛載至 docker 中，這可以使我們開發了程式碼後，即同步在 docker 內部，使開發更快速。

`entrypoint`中指定了`go run main.go`，所以我們也必須創建一個`main.go`檔，才會讓 docker 啟動時有檔案執行

```go
// main.go
package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

const (
  // 指定要連接的DB位置
	HOST     = "db"
	DATABASE = "postgres"
	USER     = "user"
	PASSWORD = "mysecretpassword"
)

func main() {
  // 連接DB
	db, err := sql.Open(
		"postgres",
		fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", HOST, USER, PASSWORD, DATABASE),
	)
	if err != nil {
		panic(err)
	}
  // 檢查連接是否成功
	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Successfully created connection to database")

  // 啟動一個簡單的http server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!!")
	})
	log.Fatal(http.ListenAndServe(":5000", nil))
}
```

在`指定要連接的DB位置`這幾行中，`HOST`為 docker 容器的名稱，這邊要注意一下並不是使用什麼`localhost`來連接，我第一次用 Docker-Compose 時這邊卡了很久 XD，而其他的欄位就是`docker-compose.yaml`中`db`裡設置的`environment`。

## 啟動 Docker-Compose

將 docker 內部需要 build 的容器都 build

```bash
$ docker-compose up
```

build 完畢後啟動

```bash
$ docker-compose up
```

![](https://i.imgur.com/N0mREyT.png)

當你看到`Successfully created connection to database`字串時就代表成功啦～

有了好環境後就可以開始開發囉，接下來下篇要介紹如何透過各式各樣的 Docs tool 來幫助我們開發更順暢！

---

謝謝你的閱讀，也歡迎分享討論指正～
