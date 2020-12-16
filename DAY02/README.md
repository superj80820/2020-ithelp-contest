本文章同時發佈於：

- [Github(包含程式碼)](https://github.com/superj80820/2020-ithelp-contest/blob/master/DAY02)
- [Medium](https://medium.com/%E9%AB%92%E6%A1%B6%E5%AD%90/day2-%E5%BE%AE%E6%9C%8D%E5%8B%99%E7%9C%9F%E7%9A%84%E6%9C%89%E9%82%A3%E9%BA%BC%E5%A5%BD-%E4%BD%A0%E7%9A%84-backend-%E6%9C%89%E9%9C%80%E8%A6%81%E9%80%B2%E5%8C%96%E6%88%90%E4%BB%96%E5%97%8E-13b801ecc7f2)
- [iT 邦幫忙](https://ithelp.ithome.com.tw/articles/10237712)

```
文章為自己的經驗與夥伴整理的內容，設計沒有標準答案，如有可以改進的地方，請告訴我，我會盡我所能的修改，謝謝大家～
```

---

近年來微服務已經被許多公司如 Google、Facebook、Netflix 採用，他到底解決了什麼單體式服務無法解決的問題呢？

## 單體式(Monolithic)的優缺點

優點:

1. 整體 Backend 設計簡單，要實踐的功能全部都在一個 Service 裡，而不用思考與其他微服務交互
2. 很好整合測試，因為測試不需要考慮與其他微服務交互，你只要針對此 Service 就好了

缺點:

1. 當邏輯開始龐大複雜起來，更改功能就變得千難萬難
2. 不能獨立 Scaling 特定功能，要 Scaling 就是所有功能一起
3. 對於語言的依賴會非常的高

## 微服務(Microservice)的優缺點

優點:

1. 功能單純，不會產生一個 Service 有`會員功能`又有`聊天功能`的現象
2. 易於 Scaling，需要強化`聊天功能`時就針對聊天 Service Scaling 即可
3. 語言的依賴極低，團隊可依照專案的需求選擇語言，只要每個微服務之間的溝通有定義好即可

缺點:

1. 整體 Backend 設計複雜，一個問題出錯你必須考慮多個微服務交互的問題
2. 整合測試困難，測試此為服務必須考量到所有相關的微服務連接

---

我們發現，其實單體式與微服務保有的優缺點剛好是相反的。

## 舉個例子

如果我們要設計一個`登入數碼世界的會員系統`，會需要什麼呢？

- 架設 Restful API Service，例如: [Go-gin](https://github.com/gin-gonic/gin)
- RDBMS 儲存會員資料，例如: Postgresql
- Cache System 來儲存 Session，例如: Redis

我們很自然而然的就會將 Service 如下圖進行 MVC 的分層，

[//]: # "(./02-01.drawio.png)"

![](https://i.imgur.com/Y21ZITn.png)

但這個時候如果要加入一個`在登入頁面會顯示的真人小窗幫手`呢？

![](https://i.imgur.com/EpYF0lT.png)

我們可能需要:

- Websocket 的聊天 Service，例如: [Go-socket.io](https://github.com/googollee/go-socket.io)
- 用來傳送分佈式訊息的 MQ System，例如: [Go-Nats](https://github.com/nats-io/nats.go)

所以 Backend 的結構可能會變成下圖:

[//]: # "(./02-02.drawio.png)"

![](https://i.imgur.com/q8REm9f.png)

這導致我們要將兩者的 Backend 結構合成在一起，這其實不會很難設計(`單體式優點1`)，也很好整合測試(`單體式優點2`)，因為只需要啟動一個 Service 就可以測試了。

[//]: # "(./02-03.drawio.png)"

![](https://i.imgur.com/RzsOTa0.png)

但可以發現，Restful API 是一進一出，而 Websocket 是用 listener 來監聽此連線的各種 event，兩者的呼叫方式不同。與外部資源的相依也不同，但卻都放在同一個 Service 中。

這導致改功能的風險是較大的(`單體式缺點1`)，客服大爆滿也無法針對 Websocket 功能 Scaling(`單體式缺點2`)。

假設這個 Service 是採用 C#，Backend 團隊在設計真人小窗幫手的 Websocket 系統想使用 Golang 這個高併發小能手，那將是不可能的任務，因為沒辦法把 C#與 Golang 組在一起寫 Service(`單體式缺點3`)。

---

微服務恰好把上面最彆扭的地方解決，就是把 Restful API 與 Websocket 拆成兩個服務。

[//]: # "(./02-04.drawio.png)"

![](https://i.imgur.com/bg1S4Ic.png)

這使 Service 的功能單純(`微服務優點1`)，一個是「會員系統」一個是「聊天系統」。

要單純針對 Websocket Service Scaling 也是沒問題的(`微服務優點2`)

此兩個 Service 用不同語言也沒問題(`微服務優點3`)，因為彼此實作並沒有相關，我們微服務有確實連接即可(`紅色框框`標起來的部分)。

But! 就是這個 But，`紅色框框`正是微服務一個大大的問題。

Websocket Service 有可能呼叫 Restful API Service 因為網路問題而產生錯誤，我們必須考慮到當 Resful API Service 掛點時要如何啟動另一個新 Restful API Service，還要考慮到 Websocket Service 要如何知道新 Service 的位置(即`Service Discover`)，這導致`微服務缺點1`。

而要整合測試，也必須把所有微服務連接都連接好，當你的微服務有十幾百個，這將是一個大工程(`微服務缺點2`)。

## 你需要微服務嗎？

我自己是認為:

單體式適合解決單純的架構，如果你負責的後端架構並不是很複雜，其實可以直接採用單體式服務即可。

微服務適合解決複雜的架構，如果有 Scaling 需求，並且分割成不同 Service 來獨立功能，那就會較適合微服務。

---

謝謝你的閱讀，也歡迎分享討論指正～

## 參考與引用資料

[Microservices vs Monolith: which architecture is the best choice for your business?](https://www.n-ix.com/microservices-vs-monolith-which-architecture-best-choice-your-business/)
