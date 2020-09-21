本文章同時發佈於：

- [Github(包含程式碼)](https://github.com/superj80820/2020-ithelp-contest/blob/master/DAY07)
- [Medium](https://medium.com/%E9%AB%92%E6%A1%B6%E5%AD%90/day8-%E8%AE%93%E4%BD%A0%E7%9A%84-backend-%E8%90%AC%E7%89%A9%E7%9A%86%E8%99%9B-%E8%90%AC%E4%BA%8B%E7%9A%86%E5%8F%AF%E6%B8%AC-clean-architecture-%E6%B8%AC%E8%A9%A6%E7%AF%87-324fddd7db9a)
- [iT 邦幫忙](https://ithelp.ithome.com.tw/articles/10241698)

```
文章為自己的經驗與夥伴整理的內容，設計沒有標準答案，如有可以改進的地方，請告訴我，我會盡我所能的修改，謝謝大家～
```

---

大家好，繼昨天[DAY07](https://github.com/superj80820/2020-ithelp-contest/blob/master/DAY07)的介紹後，Clean Architecture 的威力大家已經見識到了，我們有著`獨立可替換的彈性架構`，那還有沒有什麼優點呢？有的，就是:

> 高可測試(Testability)性

## 一個簡單的測試

你可能像我一樣曾經寫過這樣的程式碼:

```Go
// ... 其他程式碼

import "service"

func GetName() int {
  nameService := service.NewNameService()
  result := nameService.GetName()
  // ... 關於result的演算
  return result
}

// ... 其他程式碼
```

如果有天要 unit test `result的演算`，必須要使`nameService.GetName()`會回傳固定值，而不是真的打到外部 API。

因為這樣如果測試出錯，我們才能精確的說是演算錯誤了，而非外部 API 出事了。

使`nameService.GetName()`回傳固定值的做法稱為`mock`，就是將`nameService.GetName()`替換成一個`會回傳固定值的替身`。

但剛剛得程式碼，我們根本沒有機會進行替換動作，因為`nameService`在`GetName`裡頭產生，其實我們只要做一個簡單的小動作，就可以避免這個事情發生，就是:

> 將所需的變數、實體、函示從外部帶入

```Go
package logic

import "service"

func GetName(nameService *service.Engine) int {
  result := nameService.GetName()
  // ... 關於result的演算
  return result
}

// ... 其他程式碼
```

那麼我們就可以從外部做一個假的`nameService`丟入了，

```Go
// ... 其他程式碼

import (
  "logic"
  "mock"
  "github.com/stretchr/testify/assert"
)

func TestGetName(t *testing.T) {
  mockNameService := mock.NewNameService()
  result := logic.GetName(&mockNameService)
  assert.Equal(t, 3, result)
}
```

摁...這不就單純只是不要把東西寫死在 function 中嗎？

對，其實這就是依賴注入(DI)的精神之一，只不過為了要避免爆炸，所以又需要 interface 去告訴不同 caller，此參數有哪些 function，這導致其他的理論產生。

## 如果是這樣的話那我不就每層都可以測試了？！

是的，因為:

> Clean Architecture 每層都用依賴注入(DI)，所以你每層都可以換成 Mock 實體來測試

接下來我們將說明每層測試的`重點`與`第一次測試可能會困惑的地方`。

以下的 code 全部都在[Github-DAY07](https://github.com/superj80820/2020-ithelp-contest/blob/master/DAY07)，可以直接 clone 下來參考對照會比較有感覺。

### Repository 層

```Go
// ... 其他程式碼

import (
	"context"
	"testing"

	digimonPostgresqlRepo "go-server/digimon/repository/postgresql"

	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"

	"go-server/domain"
)

func TestGetByID(t *testing.T) {
  // !!! 講解1 !!!
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

  // !!! 講解2 !!!
	mockDigimon := &domain.Digimon{
		ID:     "69770f2d-933e-474d-8357-a2f8a9c874df",
		Name:   "Customer",
		Status: "Good",
	}

  // !!! 講解3 !!!
	rows := sqlmock.NewRows([]string{"id", "name", "status"}).
		AddRow(mockDigimon.ID, mockDigimon.Name, mockDigimon.Status)

  // !!! 講解4 !!!
	query := "SELECT id, name, status FROM digimons WHERE id =?"

  // !!! 講解4 !!!
  mock.ExpectQuery(query).WithArgs("69770f2d-933e-474d-8357-a2f8a9c874df").WillReturnRows(rows)
  // !!! 講解5 !!!
	d := digimonPostgresqlRepo.NewpostgresqlDigimonRepository(db)
  aDeit, _ := d.GetByID(context.TODO(), "69770f2d-933e-474d-8357-a2f8a9c874df")
  // !!! 講解5 !!!
	assert.Equal(t, mockDigimon, aDeit)
}

// ... 其他程式碼
```

> 測試重點: 呼叫外部物件的方式是否正確，並非測試外部物件

- `講解1 - go-sqlmock製作一個假的mock DB`: 如重點所說，我們並不關心外部 DB 的實際狀況，我們只關心`組出來的SQL command`是否正確，所以可以用 go-sqlmock 做一個假的 DB。
- `講解2 - 產生在mock DB需要的假資料`: 既然要測試的是`GetByID`的邏輯，那就必須先產生一組假的資料，讓`GetByID`真的有資料可以撈取。
- `講解3 - 將假資料轉換成row的struct`: `講解2`產生的資料是`domain.Digimon`的 data struct，我們必須把它轉換成 row 的 data struct，才會符合真實對 DB 取 row 的狀況。
- `講解4 - 預定mock DB會接收到何種SQL command`: 這邊是重點，我們還必須預定好會回傳的值，mock DB 在接收到正確的 command 就會回傳預定的回傳值。
- `講解5 - 把 mock DB 丟入後 Repository 層，驗證結果是否正確`

第一次測試時你可能會困惑：「我都把 DB 做成假的了，那我在測試什麼？」，答案是`SQL command是否組對`。

不論你用 gorm, SQL command，當傳進來的參數進行一連串運算後，最後一定會變為一個 SQL command，所以先設定好 mock DB`要接受什麼樣的command`與`接收正確之後要回傳什麼`是測試目的，並且最後可以用回傳的值來驗證是否是我們一開始設定好的來驗證。

### Usecase 層

```Go
// ... 其他程式碼

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"errors"
	ucase "go-server/digimon/usecase"
	"go-server/domain"
	"go-server/domain/mocks"
)

func TestGetByID(t *testing.T) {
  // !!! 講解1 !!!
	mockDigimonRepo := new(mocks.DigimonRepository)
	mockDigimon := domain.Digimon{
		ID:     "2e72c27e-0feb-44e1-89ef-3d58fd30a1b3",
		Name:   "Metrics",
		Status: "Good",
	}

  // !!! 講解2 !!!
	t.Run("Success", func(t *testing.T) {
    // !!! 講解3 !!!
		mockDigimonRepo.
			On("GetByID", mock.Anything, mock.MatchedBy(func(value string) bool { return value == "e9addf2d-8739-427a-8b30-2383b9b045b1" })).
			Return(&mockDigimon, nil).Once()

    // !!! 講解4 !!!
		u := ucase.NewDigimonUsecase(mockDigimonRepo)
		aDigimon, err := u.GetByID(context.TODO(), "e9addf2d-8739-427a-8b30-2383b9b045b1")

    // !!! 講解5 !!!
		assert.NoError(t, err)
		assert.NotNil(t, aDigimon)

		mockDigimonRepo.AssertExpectations(t)
	})
	t.Run("Fail", func(t *testing.T) {
		mockDigimonRepo.
			On("GetByID", mock.Anything, mock.MatchedBy(func(value string) bool { return value == "e9addf2d-8739-427a-8b30-2383b9b045b1" })).
			Return(nil, errors.New("Get error")).Once()

		u := ucase.NewDigimonUsecase(mockDigimonRepo)
		aDigimon, err := u.GetByID(context.TODO(), "e9addf2d-8739-427a-8b30-2383b9b045b1")

		assert.Error(t, err)
		assert.Nil(t, aDigimon)

		mockDigimonRepo.AssertExpectations(t)
	})
}

// ... 其他程式碼
```

> 測試重點: 業務邏輯路線是否預期，並非測試 Repository 層

Usecase 層沒有像 go-sqlmock 這樣通用的 mock 套件可以使用，但是還記得有 Domain 層定好的 interface 嗎？

我們可以使用[mockery](https://github.com/vektra/mockery)，`依照定好的interface生產出mock物件`，依照[mockery](https://github.com/vektra/mockery)安裝好後:

```bash
$ cd go-server/domain
$ mockery --all --keeptree
```

![](https://i.imgur.com/LOx4BCl.png)

此時會多出 mocks 資料夾，這就是`mock data struct`

- `講解1 - 透過mocks資料夾產生mockDigimonRepo`
- `講解2 - 依照不同情境區分測試`: Usecase 層比較會有 `不同的程式路線`產生，例如：「數碼蛋參數如果沒帶的話、如果 DB 爆掉的話等等」，所以我們可以依照這些情境來區分測試
- `講解3 - 設定mockDigimonRepo預定接收的參數與回傳的資料`: 這邊是最有趣的部分，你可以透過`.On()`來定義 mockDigimonRepo 在 Usecase 層裡被呼叫的時候預定要接收什麼，並且回傳資料供邏輯繼續使用。
- `講解4 - 將mockDigimonRepo丟入Usecase層並實際運行`。
- `講解5 - 驗證Usecase回傳資料與mockDigimonRepo運行的結果`: `.AssertExpectations`要稍微注意一下，他是在驗證`.On()`是否真的有被 call 到。

在`講解3`處，有著強大的功能:

- `.On()`: 預期要呼叫什麼 function。
- `mock.MatchedBy()`: 預期參數要長什麼樣子。
- `.Return()`: 指定回傳值。
- `.Once()`: 將此`.On`效果只作用一次，這可以使你設定`同個function`在`不同次的.On`都有不同結果。
- `.Run()`: 雖然沒有出現在此測試，但非常有用。Golang 有者許多`bind(&body)`的指標綁定資料方式，這不是一進一出的方式，所以無法用`.Return()`來指定結果。而`.Run()`可以捕捉`bind(&body)`中的`body`指標位置，並且修改其內容，以達到指定結果的效果

這使業務邏輯`所有的路線都可測試到`，真的是非常有趣 XD。

### Delivery 層

```Go
import (
	"bytes"
	"encoding/json"
	"go-server/domain"
	"go-server/domain/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	digimonHandlerHttpDelivery "go-server/digimon/delivery/http"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetDigimonByDigimonID(t *testing.T) {
	mockDigimon := domain.Digimon{
		ID:     "8c862535-6de2-4da2-ad21-c853e4343bd7",
		Name:   "III",
		Status: "good",
	}
	mockDigimonMarshal, _ := json.Marshal(mockDigimon)
	mockDigimonCase := new(mocks.DigimonUsecase)

	mockDigimonCase.On("GetByID", mock.Anything, mockDigimon.ID).Return(&mockDigimon, nil)

	r := gin.Default()

	handler := digimonHandlerHttpDelivery.DigimonHandler{
		DigimonUsecase: mockDigimonCase,
	}

	r.GET("/api/v1/digimons/:digimonID", handler.GetDigimonByDigimonID)

  // !!! 講解1 !!!
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/digimons/8c862535-6de2-4da2-ad21-c853e4343bd7", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, string(mockDigimonMarshal), w.Body.String())
}
```

> 測試重點: Delivery 層是否都正確接收了引擎的參數，並非測試 Usecase 層

Delivery 層與 Usecase 層概念相同(事實上每層都相同)，都是創造上一層的 mock 實體，並專注測試此層的邏輯。

要注意的地方是，此測試並非真的起了一個 Server 以外部打 API 近來的方式測試。

而是在`講解1`處之前設定好了 router 後，透過`httptest.NewRecorder()`產生一個 request，以`r.ServeHTTP(w, req)`的方式丟入 Golang-Gin 的引擎中。

## 好用的測試指令

一次跑全部專案的測試，你可以把此指令放上 CI，這樣就可以在每次 CD 之前來測試一下！

```
$ cd go-server
$ go test ./...
```

![](https://i.imgur.com/RnuAZbr.png)

(`cmd/main.go:31:3`可以不必理會，那是因為讀取`.env`失敗的問題，但跟測試無關)

但是否我們程式中有`沒測到的路線`呢？這也是可以靠 Golang 自動偵測的，運行以下指令:

```bash
$ cd go-server
$ go test -coverprofile cover.out ./...
$ go tool cover -html=cover.out -o cover.html
$ open cover.html
```

最後會開啟此 html file，

![](https://i.imgur.com/7RJSRWG.png)

可以看到上方可以顯示測試覆蓋率(就是路線覆蓋率)，而下方會顯示沒測試到的 code，太神奇啦！

## 參考

- [bxcodec/go-clean-arch](https://github.com/bxcodec/go-clean-arch)
- [Trying Clean Architecture on Golang](https://medium.com/hackernoon/golang-clean-archithecture-efd6d7c43047)
- [unit testing - How to measure code coverage in Golang? - Stack Overflow](https://stackoverflow.com/questions/10516662/how-to-measure-code-coverage-in-golang)
- [testing - How to `go test` all tests in my project? - Stack Overflow](https://stackoverflow.com/questions/16353016/how-to-go-test-all-tests-in-my-project)
- [unit testing - Mock interface method twice with different input and output using testify - Stack Overflow](https://stackoverflow.com/questions/62165773/mock-interface-method-twice-with-different-input-and-output-using-testify)
- [Using testify for Golang tests – ncona.com – Learning about computers](https://ncona.com/2020/02/using-testify-for-golang-tests/)
