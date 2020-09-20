本文章同時發佈於：

- [Github(包含程式碼)]()
- [Medium]()
- [iT 邦幫忙]()

```
文章為自己的經驗與夥伴整理的內容，設計沒有標準答案，如有可以改進的地方，請告訴我，我會盡我所能的修改，謝謝大家～
```

---

大家好，繼昨天[DAY05](https://github.com/superj80820/2020-ithelp-contest/blob/master/DAY05) Docs tool 的介紹後，我們要利用這些 Docs 來產生 Golang Server 介面，並以 Clean Architecture 實作。

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

### Controller/ 層多了 Usecase 層

> 業務邏輯放置 Usecase 層。而 Controller/Delivery 只負責交代如何把這些 Usecase 帶給 View 層。

如果沒有 Usecase 層，常常會導致 Controller 層的邏輯多到滿出來，我常碰到的問題是 Controller 其實也是有許多邏輯適合分享，但是傳統 MVC 卻沒有方式去傳遞他，現在有 Usecase 層，就會使 Controller 層就清爽許多。

> 在沒有任何業務邏輯下 Controller 層更加專心負責`交付`，因此稱之為 Delivery 層會更加合適。

### View 層變得多樣且彈性

由於 Delivery 層負責交付 Usecase 給 View 層，所以

> 要更換任何 View 層只需要重寫 Delivery 層，業務邏輯的 Usecase 是不需要改的

這實在是非常棒！某次要從 Restful API 更換成 Websocket 介面時，我直接

![](https://i.imgur.com/F1iY76L.png)

再回頭看看某些專案，只有 Controller 層導致業務邏輯已經完全黏在 View 層，我只想大喊：「你為什麼不早說 QQ」。

順帶一提，MVC 的 View 其實最初是真的要有`看得到的網頁頁面`的情境，因為以前的 Server 都是直接吐網站的，不過現在 Server 的設計已經跟過往大不相同，所以我喜歡把 View 稱為`要呈現給Client端的方法`，所以 Restful API, Websocket, gRPC 我都認為算是一種 View。

### 給每層建立介面的 Domain 層

除了 Delivery 層可以替換，Repository 層也是可以替換的，比如說 RDBMS 要從 MySQL 換成 PostgreSQL。

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

## 為什麼更換 DB 引擎，只要介面不變就不會爆炸？

如果你是從 Node.js 轉到 Golang，你可能會對這個概念感到陌生(就像我 XD)，以下舉一個 Node.js 大家常會遇到的困難，

```javascript
function callAPI(caller) {
  caller.get("https://api");
}
```

這時候會有一個問題：

> 傳進來的 caller 真的有 get 這個 method 嗎？

為了驗證我們會直接跑起來，

```javascript
callAPI(ACaller);
// 因為沒有get這個method
// 會顯示`Uncaught TypeError: caller.get is not a function`

callAPI(BCaller);
// 因為有get這個method，所以運作正常
```

但我們真的只能`跑起來`才能驗證`有沒有這個method`嗎？

答案是不用的，那就是透過程式的靜態分析，在程式還沒跑起來前分析 type(型別)或者是 interface(介面)適合符合預期，所以剛剛的 callAPI 可以這樣加入 interface 給參數

```javascript
// interfaceA會檢查此參數是否有get method
function callAPI(caller: interfaceA) {
  caller.get("https://api");
}
// 如果a有get method就不會噴錯，而沒有就會噴錯
callAPI(a);
```

回到原本不會爆炸的問題，我們可以替 DB 引擎`包上一層interface`，再用 JavaScript 來舉例:

```javascript
// 此code不能運行，為舉例的pseudocode
interfaceA {
  insert()
}
function insertRow(db: interface) {
  db.insert("foo");
}
const postgresDB = {
  insert(value) {
    // 這裡才是真的DB引擎的insert，我們為他包裝了一層interface，使呼叫方不會爆炸
    postgres.insert(value)
  }
}
const mysqlDB = {
  insert(value) {
    // DB引擎換了，但是外層interface與postgresDB相同
    mysqlDB.insert(value)
  }
}

// 兩者都可以運行，因為都符合interfaceA的定義
insertRow(postgresDB)
insertRow(mysqlDB)
```

所以，Golang 因為有 interface 來跟使用的的程式說該要有哪些 method，更換 DB 引擎就不會爆炸了。

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
