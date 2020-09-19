本文章同時發佈於：

- [Github(包含程式碼)](https://github.com/superj80820/2020-ithelp-contest/blob/master/DAY05)
- [Medium](https://medium.com/%E9%AB%92%E6%A1%B6%E5%AD%90/day5-%E4%BB%80%E9%BA%BC-%E4%BD%A0%E7%9A%84%E7%A8%8B%E5%BC%8F%E7%A2%BC%E7%94%B1%E6%96%87%E4%BB%B6%E7%94%A2%E7%94%9F-%E9%80%99%E6%A8%A3%E4%B8%8D%E5%B0%B1%E4%B8%8D%E7%94%A8%E8%A3%9C%E6%96%87%E4%BB%B6%E4%BA%86%E5%97%8E-docs-tool-%E7%AF%87-3e9a734d063e)
- [iT 邦幫忙](https://ithelp.ithome.com.tw/articles/10240175)

```
文章為自己的經驗與夥伴整理的內容，設計沒有標準答案，如有可以改進的地方，請告訴我，我會盡我所能的修改，謝謝大家～
```

---

大家好，繼昨天[DAY04](https://github.com/superj80820/2020-ithelp-contest/blob/master/DAY04) Docker-Compose 的介紹後，我們終於有了一個開發環境，接下來要實作夠過各種 Docs tool 來建構各種文件與程式介面。

## 為什麼需要這些 Docs tool

> 文件與實作同步，為不同團隊提供一個簡潔的系統介紹，而不是遇到任何問題就`看code`

以前接受了許多專案，文檔提供不完全，導致許多 API 根本不清楚作用，而是要透過`看code`來了解整體情境，這花費了巨量的時間。

也不是說看 code 不好，在[Clean Code](https://www.tenlong.com.tw/products/9789862017050)書中也有提到 :

> code 必須要清楚表達意圖，而不是靠各種註解解決

而文件跟註解有幾分相似，是否如果 code 寫得好就不需要文件了呢？

我認為還是要的，以下是我覺得需要寫成文件的部分:

- 一連串不同 API 達成的情境: 因為單純看 code 是無法說明一整個情境的。
- API 所規範的介面: 各團隊可以根據此規範來串接 API，以避免後端說要帶 a 資料前端卻帶 b 資料的窘境。

## 為什麼文件往往趕不上實作？

> 因為寫文件的成本往往高於寫 code

我認為有以下關係:

1. 當更新了 code 的功能，我們就要點開各式各樣的文件去更新
2. 文件軟體用起來有夠不順手，奇怪我是程式設計的為什麼要一天到晚畫圖
3. 為什麼文件寫的跟 code 不一樣，算了這個文件不看白不看，還是看 code

所以我已以下原則來寫文件，我認為可以減低以上問題的成本 XD:

1. 統一都用 Markdown 來寫，流程圖與關係圖用[mermaid-js](https://github.com/mermaid-js/mermaid)與[vscode-drawio](https://github.com/hediet/vscode-drawio)一次解決
2. 因為採用 Markdown 所以不用再`為了邏輯而畫圖`，而是寫好邏輯就自動產生圖
3. 用[swagger-generator](https://github.com/swagger-api/swagger-codegen)來產生程式介面，再來開發，這樣就可以`用文件產生程式介面`。

## 用文件產生程式介面？！

沒錯！你可以這樣做，在 GraphQL 問世後，大家慢慢發現可以這樣做，要更改程式嗎？那請先更新 GraphQL schema，如此一來文件永遠不會落後，前後端的 code 也可以透過這個文件產生。

而 gRPC 也是有同樣的概念在裡頭，gRPC 是透過定義好 Protobuf 文件來產生程式碼。

但較少人知道，Restful API 其實也是可以這樣做的，即是定義好 Open API，並透過上述[swagger-generator](https://github.com/swagger-api/swagger-codegen)來產生程式碼。

![](https://i.imgur.com/KI8pd9e.jpg)

## Docs tool 介紹

- [mermaid-js](https://github.com/mermaid-js/mermaid): 提供用簡單的 Markdown 來撰寫流程圖的功能
- [vscode-drawio](https://github.com/hediet/vscode-drawio): 可以讓 vscode 可以直接在`.jpg/.png`中編輯圖片並同步更新在 Markdown 中
- [swagger-generator](https://github.com/swagger-api/swagger-codegen): [swagger](https://github.com/swagger-api)是一個撰寫 Open API 介面的方案，大家常常拿它來撰寫介面，但其實他也可以透過這些文件來生產出`不同語言的Server端與Client端`，對於規範好不同團隊語言的接口很方便。
- [Insomnia Designer](https://insomnia.rest/products/designer/): 可以用 Open API 來產生 Postman 測試端點的軟體！

## 開始來基於 Digimon-Service 來撰寫文件～

[Example Github code 在此](https://github.com/superj80820/2020-ithelp-contest/blob/master/DAY05)

---

![From Youtube: 2019 Bandai 20th Anniversary Digimon Virtual Pet Unboxing & Review [4K UHD]](https://i.imgur.com/mtnMhIk.jpg)

還記得小時候玩的數碼育成機嗎？一開始會有一顆數碼蛋，我們要慢慢培育他，隨著時間這顆數碼蛋就會孵化出數碼獸，並與我們一起成長。Digimon-Service 希望可以提供這些功能。

所以基於這些需求，我們開始把文件撰寫出來～

---

Digimon-Service 結構圖 - 使用[vscode-drawio](https://github.com/hediet/vscode-drawio)完成:

[//]: # "(./digimon-service.drawio.png)"

![](https://i.imgur.com/qgvVX2C.png)

提供的功能：

- 提供一個數碼蛋，讓 client 端可以領取。
- 數碼蛋會隨著時間孵化。

---

實作我們採用 Restful API，以下是 API - 使用[Insomnia Designer](https://insomnia.rest/products/designer/)完成:

```yaml
# swagger.yaml
openapi: 3.0.1
info:
  title: Digimon Service API
  description: 提供孵化數碼蛋與培育等數碼寶貝養成服務
  version: 1.0.0
servers:
  - url: http://localhost:5000/api/v1
paths:
  /digimons:
    post:
      summary: 產生數碼蛋
      description: 產生一顆數碼蛋，供request端養成
      requestBody:
        description: 客製數碼蛋的請求
        required: true
        content:
          application/json:
            schema:
              $ref: "#/components/schemas/DigimonInfoRequest"
      responses:
        "200":
          description: 數碼蛋的資訊
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/DigimonInfo"
        "500":
          $ref: "#/components/responses/500InternalError"
  /digimons/{digimonID}:
    get:
      summary: 查看數碼獸狀態
      parameters:
        - in: path
          name: digimonID
          schema:
            type: string
          required: true
          description: 數碼蛋的唯一識別碼，格式為uuid v4
      responses:
        "200":
          description: 數碼蛋的資訊
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/DigimonInfo"
        "500":
          $ref: "#/components/responses/500InternalError"
  /digimons/{digimonID}:foster:
    post:
      summary: 培育數碼獸
      description: 對數碼獸進行培育，以改善數碼獸的狀態
      parameters:
        - in: path
          name: digimonID
          schema:
            type: string
          required: true
          description: 數碼蛋的唯一識別碼，格式為uuid v4
      requestBody:
        description: 培育的食物
        required: true
        content:
          application/json:
            schema:
              $ref: "#/components/schemas/FosterRequest"
      responses:
        "200":
          description: 培育完畢後的數碼獸的資訊
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/DigimonInfo"
        "500":
          $ref: "#/components/responses/500InternalError"
components:
  schemas:
    DigimonInfo:
      type: object
      properties:
        id:
          type: string
          description: 數碼蛋的唯一識別碼，格式為uuid v4
          example: 56e400bd-c98b-49b9-ad8c-0607800e026f
        name:
          type: string
          description: 數碼蛋的名稱
          example: Agumon
        status:
          type: string
          description: 數碼蛋此時的狀態
          example: healthy
    FosterRequest:
      type: object
      properties:
        food:
          type: object
          description: 培育所使用的食物
          properties:
            name:
              type: string
              description: 食物名稱
              example: apple
    DigimonInfoRequest:
      type: object
      properties:
        name:
          type: string
          description: 數碼蛋的名字
      required:
        - name
      example:
        name: Agumon
    Error:
      type: object
      properties:
        message:
          type: string
          description: 錯誤訊息
        code:
          type: number
          description: >
            錯誤代碼:
             * `3000` - Internal error
      required:
        - message
        - code
      example:
        message: "Internal error. Parsing failed"
        code: 3000
  responses:
    500InternalError:
      description: 伺服器錯誤
      content:
        application/json:
          schema:
            $ref: "#/components/schemas/Error"
```

[Insomnia Designer](https://insomnia.rest/products/designer/)的設計介面非常友好，並且！當你點了上方的`DEBUG`

![](https://i.imgur.com/GZT8BZz.png)

他竟人產生了測試端點！連端點都不用慢慢加了，太神啦～！

![](https://i.imgur.com/mmrEOXg.png)

---

資料庫如下 - 使用[vscode-drawio](https://github.com/hediet/vscode-drawio)完成:

[//]: # "(./er.drawio.png)"

![](https://i.imgur.com/TUI9Sr5.png)

- Digimons: 紀錄數碼獸狀態
- Diets: 紀錄數碼獸飲食

---

流程上如下 - 使用[mermaid-js](https://github.com/mermaid-js/mermaid)完成:

產生數碼蛋

[![](https://mermaid.ink/img/eyJjb2RlIjoiXG5zZXF1ZW5jZURpYWdyYW1cbiAgICBwYXJ0aWNpcGFudCBjIGFzIENsaWVudFxuICAgIHBhcnRpY2lwYW50IHMgYXMgRGlnaW1vbi1TZXJ2aWNlXG4gICAgcGFydGljaXBhbnQgZGIgYXMgUG9zdGdyZVNRTFxuICAgIGMtPj5zOiDkvb_nlKggQVBJIFBPU1QgYXBpL3YxL2RpZ2ltb25zXG4gICAgcy0-PnM6IOeUoueUn-aVuOeivOibi--8jOWNs-aYr0RpZ2ltb25zIHRhYmxl6ZyA6KaB55qE6LOH5paZXG4gICAgcy0-PmRiOiDlsIfmlbjnorzom4vlrZjpgLJEaWdpbW9ucyB0YWJsZVxuICAgIHMtLT4-Yzog5Zue5YKz5oiQ5YqfXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9fQ)](https://mermaid-js.github.io/mermaid-live-editor/#/edit/eyJjb2RlIjoiXG5zZXF1ZW5jZURpYWdyYW1cbiAgICBwYXJ0aWNpcGFudCBjIGFzIENsaWVudFxuICAgIHBhcnRpY2lwYW50IHMgYXMgRGlnaW1vbi1TZXJ2aWNlXG4gICAgcGFydGljaXBhbnQgZGIgYXMgUG9zdGdyZVNRTFxuICAgIGMtPj5zOiDkvb_nlKggQVBJIFBPU1QgYXBpL3YxL2RpZ2ltb25zXG4gICAgcy0-PnM6IOeUoueUn-aVuOeivOibi--8jOWNs-aYr0RpZ2ltb25zIHRhYmxl6ZyA6KaB55qE6LOH5paZXG4gICAgcy0-PmRiOiDlsIfmlbjnorzom4vlrZjpgLJEaWdpbW9ucyB0YWJsZVxuICAgIHMtLT4-Yzog5Zue5YKz5oiQ5YqfXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9fQ)

培育數碼獸

[![](https://mermaid.ink/img/eyJjb2RlIjoiXG5zZXF1ZW5jZURpYWdyYW1cbiAgICBwYXJ0aWNpcGFudCBjIGFzIENsaWVudFxuICAgIHBhcnRpY2lwYW50IHMgYXMgRGlnaW1vbi1TZXJ2aWNlXG4gICAgcGFydGljaXBhbnQgZGIgYXMgUG9zdGdyZVNRTFxuICAgIGMtPj5zOiDkvb_nlKggQVBJIFBPU1QgYXBpL3YxL2RpZ2ltb25zLzpkaWdpbW9uSUQvZm9zdGVyXG4gICAgcy0-PmRiOiDlsIdyZXF1ZXN055qEZGlldOizh-aWmeWtmOWFpURpZXRzIHRhYmxlXG4gICAgcy0-PmRiOiDmm7TmlrBEaWdpbW9uIHRhYmxl55qEc3RhdHVz6Iez5YGl5bq3XG4gICAgcy0tPj5jOiDlm57lgrPmiJDlip9cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In19)](https://mermaid-js.github.io/mermaid-live-editor/#/edit/eyJjb2RlIjoiXG5zZXF1ZW5jZURpYWdyYW1cbiAgICBwYXJ0aWNpcGFudCBjIGFzIENsaWVudFxuICAgIHBhcnRpY2lwYW50IHMgYXMgRGlnaW1vbi1TZXJ2aWNlXG4gICAgcGFydGljaXBhbnQgZGIgYXMgUG9zdGdyZVNRTFxuICAgIGMtPj5zOiDkvb_nlKggQVBJIFBPU1QgYXBpL3YxL2RpZ2ltb25zLzpkaWdpbW9uSUQvZm9zdGVyXG4gICAgcy0-PmRiOiDlsIdyZXF1ZXN055qEZGlldOizh-aWmeWtmOWFpURpZXRzIHRhYmxlXG4gICAgcy0-PmRiOiDmm7TmlrBEaWdpbW9uIHRhYmxl55qEc3RhdHVz6Iez5YGl5bq3XG4gICAgcy0tPj5jOiDlm57lgrPmiJDlip9cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In19)

查看數碼獸

[![](https://mermaid.ink/img/eyJjb2RlIjoiXG5zZXF1ZW5jZURpYWdyYW1cbiAgICBwYXJ0aWNpcGFudCBjIGFzIENsaWVudFxuICAgIHBhcnRpY2lwYW50IHMgYXMgRGlnaW1vbi1TZXJ2aWNlXG4gICAgcGFydGljaXBhbnQgZGIgYXMgUG9zdGdyZVNRTFxuICAgIGMtPj5zOiDkvb_nlKggQVBJIEdFVCBhcGkvdjEvZGlnaW1vbnMvOmRpZ2ltb25JRFxuICAgIHMtPj5kYjog5pKI5Y-WRGlnaW1vbiB0YWJsZeeahOaVuOeivOeNuOizh-aWmVxuICAgIHMtLT4-Yzog5Zue5YKz5pW456K854246LOH5paZXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9fQ)](https://mermaid-js.github.io/mermaid-live-editor/#/edit/eyJjb2RlIjoiXG5zZXF1ZW5jZURpYWdyYW1cbiAgICBwYXJ0aWNpcGFudCBjIGFzIENsaWVudFxuICAgIHBhcnRpY2lwYW50IHMgYXMgRGlnaW1vbi1TZXJ2aWNlXG4gICAgcGFydGljaXBhbnQgZGIgYXMgUG9zdGdyZVNRTFxuICAgIGMtPj5zOiDkvb_nlKggQVBJIEdFVCBhcGkvdjEvZGlnaW1vbnMvOmRpZ2ltb25JRFxuICAgIHMtPj5kYjog5pKI5Y-WRGlnaW1vbiB0YWJsZeeahOaVuOeivOeNuOizh-aWmVxuICAgIHMtLT4-Yzog5Zue5YKz5pW456K854246LOH5paZXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9fQ)

## 總結

有了這些便捷的工具後，

> 文件就不是個很`花費時間又重工`的事情了，而是可以直接`使用在開發上`的好幫手！

接下來要介紹如何透過[swagger-generator](https://github.com/swagger-api/swagger-codegen)生產出 Golang Server 的介面並且透過 Clean Architecture 來實作，並透過[Insomnia Designer](https://insomnia.rest/products/designer/)測試

---

謝謝你的閱讀，也歡迎分享討論指正～

## 參考

- [2019 Bandai 20th Anniversary Digimon Virtual Pet Unboxing & Review [4K UHD]](https://www.youtube.com/watch?v=FuzMB5y8rOw)
