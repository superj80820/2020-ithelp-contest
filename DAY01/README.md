本文章同時發佈於：

- [Github(包含程式碼)](https://github.com/superj80820/2020-ithelp-contest/blob/master/DAY01)
- [Medium](https://medium.com/%E9%AB%92%E6%A1%B6%E5%AD%90/day01-%E5%9C%A8%E9%96%8B%E5%A7%8B%E6%95%B8%E7%A2%BC%E5%BE%AE%E6%9C%8D%E5%8B%99%E4%B9%8B%E6%97%85%E5%89%8D-f00b950c4f04)
- [iT 邦幫忙](https://ithelp.ithome.com.tw/articles/10234832)

---

![](https://i.imgur.com/aLFTBaq.png)

## 前言

在微服務的世界，Golang 為什麼可以脫穎而出，而 K8s 到底幫忙了什麼，Istio 怎麼解決 Service Mesh 的？面對這琳瑯滿目的生態系，就彷彿小時後愛看的`數碼寶貝`一樣。

> 其實，不知道的事情，只要靠自己親身去體驗，然後，用自己的眼睛去看，那就好了。 - 被選召的孩子: 源輝二

秉持著這樣的踩雷模式，因此產生了此系列，再麻煩大大們指教了，謝謝～

## 為什麼有此系列

因為在工作上會越來越需要使用 K8s + Istio，在實際應用的每一天中，也把學習的內容整理成文章，希望能與大大們討論與指教～

## 系列冒險等級

此系列適合探討以下問題的讀者：

- 從單體式服務轉微服務
- Istio 與 K8s 初探
- Golang Clean Architecture 微服務的設計

如前面所說的，我並不是一個微服務的超級專家，而是一個每天碰到微服務問題的初學者，所以文章不會有太高深的問題，不過如 Huli 大大[這篇文章](https://medium.com/hulis-blog/why-blogging-ab77fd8c6ffa)所說的，也許初學能更站在站在初學者的角度來講解，也請各位大大多指教～

## 系列撰寫方向

從單體式服務(Monolithic)轉變為微服務(Microservice)，這樣的轉變確實解決了許多問題，例如：「更好的擴展性、更拆分的邏輯」，但完全不同的後端架構卻也產生了諸多問題。

例如：Backend 人員如果要處理一個情境，原本在單體式服務只是專注在一個 Service 上，但微服務所因為要考慮多個 Service 所以範圍更大，關於多個 Service 互動的愛恨情仇，大家通常把它稱為 Service Mesh，Istio 便是為了解決此問題而生。

---

所以基於以上問題，大致想分為以下三個階段來介紹：

1. Golang 微服務實作：，一個簡單自由的 Golang 微服務要怎麼用 Docs 工具規劃、Docker 包裝環境、Clean Architecture 建構。
2. K8s：當許多的 Docker Image 都完成後，就要開始部署至雲端主機上了，而 K8s 是現今最紅的 Container 管理平台，此篇幅會介紹如何對這些 Container 進行管理。
3. Istio：在許多 Container 都部署完畢後，我們發現多個微服務的溝通竟然比單體式服務複雜非常多！此篇幅會透過 Istio 來介紹怎麼解決此問題。

---

後端 Backend 的設計百百種，沒有一種是正確答案，接下來我會用我所理解的方式解釋給大家，如果大家對於文章內容也所疑問也歡迎討論～冒險要開始啦！
