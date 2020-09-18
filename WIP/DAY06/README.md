本文章同時發佈於：

- [Github(包含程式碼)]()
- [Medium]()
- [iT 邦幫忙]()

```
文章為自己的經驗與夥伴整理的內容，設計沒有標準答案，如有可以改進的地方，請告訴我，我會盡我所能的修改，謝謝大家～
```

---

[//]: # "york TODO: 新增 DAY05 連結"

大家好，繼昨天[DAY05]() Docs tool 的介紹後，我們要利用這些 Docs 來產生 Golang Server 介面，並以 Clean Architecture 實作。

## 什麼是 Clean Architecture

[The Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)是軟體界大神[Robert C. Martin](https://en.wikipedia.org/wiki/Robert_C._Martin)提出，他所寫的著名作品有[Clean Code](https://www.books.com.tw/products/0010579897), [Clean Coder](http://books.com.tw/products/0010598217)這些好書。

> 工程師︰我已經拜讀了《Clean Code》，還有必要讀《Clean Architecture》嗎？
> 架構師︰喔，你會做磚頭，那你會蓋房子嗎？

Clean Architecturer 將更面想整體的 Code 而非 Code 的單一部分。

而我使用的 Clean Architecture 結構是學習[bxcodec](https://github.com/bxcodec)的[go-clean-arch](https://github.com/bxcodec/go-clean-arch)，如圖，此結構已經經過許多人驗證，是很 nice 的結構 XD。

## Clean Architecturer 的結構

![From: bxcodec/go-clean-arch](https://i.imgur.com/c3we5K6.png)

如果看了有點矇沒關係，我們可以用傳統的 MVC 來理解他，

![](https://i.imgur.com/8Qj2ZR9.png)

### Model 層變為 Repository 層

以往單體式 Server 還不太需要與那麼多微服務溝通，所以 Model 層可以單純連接 DB 就好。

但現在其實還有很多微服務要溝通，所以我們就統一稱作`Repository`層吧！

> 凡從外部進來的事物，都可以放在`Repository`層

我常看到 Server 專案的架構是傳統 MVC，但微服務溝通這層卻不知道往哪裡擺，有些直接放在 Controller，有些又特別拉出一層，導致每個專案架構差異慎大。現在不用管了統一放在`Repository`層！

### Controller 層多了 Usecase 層

> 業務邏輯放置 Usecase 層，而 Controller 只負責交代如何把這些 Usecase 帶給 View 層

如果沒有 Usecase 層，常常會導致 Controller 層的邏輯多到滿出來，我常碰到的問題是 Controller 其實也是有許多邏輯適合分享，但是傳統 MVC 卻沒有方式去傳遞他，現在有 Usecase 層 Controller 層就清爽許多。

### View 層變得多樣且彈性

由於 Controller 層負責交付 Usecase 給 View 層，所以

> 要更換任何 View 層只需要重寫 Controller 層，業務邏輯的 Usecase 是不需要改的

這實在是非常棒！某次要從 Restful API 更換成 Websocket 介面時，我直接

![](https://i.imgur.com/F1iY76L.png)

再回頭看看某些專案，Controller 層已經完全黏在 View 層，我只想大喊：「你為什麼不早說 QQ」。

順帶一提 MVC 的 View 其實最初是真的要有`看得到的頁面`的，因為以前的 Server 都是直接吐網站的，不過現在 Server 的設計已經跟過往大不相同，所以我喜歡把 View 稱為`要呈現給Client端的方法`，所以 Restful API, Websocket, gRPC 我都認為算是 View。

### 給每層建立介面的 Domain 層

除了 Controller 層可以替換，Model 也是可以替換的，比如說 RDBMS 要從 MySQL 換成 PostgreSQL。

但要達成這個目標我們必須先訂好一個介面，比如說儲存 digimons table 可以先訂好以下介面:

```go
type Digimon struct {
  UUID string
  Name   string
  Status string
}

type DigimonRepository interface {
	Store(ctx context.Context, u *Digimon) error
}
```

更換 DB 引擎並不會造成 Usecase 層的 Call 的時候爆炸，因為介面沒變。

我們只需關心實作不同 DB 的儲存方式即可，所以這樣的特性造成:

> 每層都高獨立性的架構

## Clean Architecturer 的高可測試性

Clean Architecturer 提倡`DI - 依賴注入 (Dependency injection)`，網路有很多很仔細的介紹，但我認為可以用以下例子很簡單說明他的`精神`:

```go
❌
db := sql.connectDB()
func storeString() {
  db.store("String")
}

✅
func storeString(db *sql.DB) {
  db.store("String")
}
```

> 希望程式的實作依賴可以由外部注入，而不是被固定在底層

更白話點

> 處理的人從參數帶進來啦，不然你寫死在 function 裡我怎麼換人

有了這個概念後，如果想測試 DigimonRepository 層，你不需要真的連上 DB，而是使用運行在程式上的[mock DB](https://github.com/DATA-DOG/go-sqlmock)，並且透過 DI 替換 DB 實體:

```go
func storeString(db *sql.DB) {
  db.store("String")
}
// 用真的DB
db := sql.connectDB()
storeString(db)

// 用MockDB
db := sql.connectMockDB()
storeString(db)
```

## 參考

- [bxcodec/go-clean-arch](https://github.com/bxcodec/go-clean-arch)
- [Trying Clean Architecture on Golang](https://medium.com/hackernoon/golang-clean-archithecture-efd6d7c43047)
