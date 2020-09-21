本文章同時發佈於：

- [Github(包含程式碼)](https://github.com/superj80820/2020-ithelp-contest/blob/master/DAY07)
- [Medium](https://medium.com/%E9%AB%92%E6%A1%B6%E5%AD%90/day7-%E5%A5%94%E6%94%BE%E7%9A%84-golang-%E5%8D%BB%E9%9A%B1%E8%97%8F%E8%91%97%E6%9C%89%E7%B4%80%E5%BE%8B%E7%9A%84%E6%9E%B6%E6%A7%8B-clean-architecture-%E5%AF%A6%E4%BD%9C%E7%AF%87-dd41610dcde7)
- [iT 邦幫忙](https://ithelp.ithome.com.tw/articles/10241479)

```
文章為自己的經驗與夥伴整理的內容，設計沒有標準答案，如有可以改進的地方，請告訴我，我會盡我所能的修改，謝謝大家～
```

---

大家好，繼昨天[DAY06](https://github.com/superj80820/2020-ithelp-contest/blob/master/DAY06)的介紹後，相信大家已經對 Clean Architecture 的稍有概念了，接下來將介紹實作的部分，相信會讓各位更理解 Clean Architecture 的好處。

不過，你也可以先進入 DAY07 的資料夾底下把 Server 運行起來，這樣會比較有感覺。

```bash
$ cd DAY07
$ docker-compose up
```

## 透過[swagger-generator](https://github.com/swagger-api/swagger-codegen)來產生 Server 介面

進到[DAY07](https://github.com/superj80820/2020-ithelp-contest/blob/master/DAY07)的資料夾，使用 docker 運行以下指令:

```bash
$ docker run --rm -v ${PWD}:/local swaggerapi/swagger-codegen-cli-v3 generate \
    -i /local/swagger.yaml \
    -l go-server \
    -o /local/go-server
```

swagger-generator 會自動產生以下的 code:

![](https://i.imgur.com/pnEyUT8.png)

接下來我們將一步一步實作成以下的 code:

![](https://i.imgur.com/DcBLIhV.png)

## 來實作吧！

![](https://i.imgur.com/8Qj2ZR9.png)

![](https://i.imgur.com/xDYVMWG.png)

請配合此兩張圖與並依照[Github 範例](https://github.com/superj80820/2020-ithelp-contest/blob/master/DAY07)一步一步實作，

### Domain 層 - 規範一切的老大哥

如[DAY06](https://github.com/superj80820/2020-ithelp-contest/blob/master/DAY06)所說，

> 我們需要一個`interface`來告訴每個 call 的程式他們到底在 call 什麼

所以必須建立建立 diet(培育)與 digimon(數碼獸)兩個 interface 在 domain 層。

![](https://i.imgur.com/HOkkLhs.png)

以`digimon.go`來說，

```golang
package domain

import "context"

// !!! 講解1 !!!
// Digimon ...
type Digimon struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

// !!! 講解2 !!!
// DigimonRepository ...
type DigimonRepository interface {
	GetByID(ctx context.Context, id string) (*Digimon, error)
	Store(ctx context.Context, d *Digimon) error
	UpdateStatus(ctx context.Context, d *Digimon) error
}

// !!! 講解3 !!!
// DigimonUsecase ..
type DigimonUsecase interface {
	GetByID(ctx context.Context, id string) (*Digimon, error)
	Store(ctx context.Context, d *Digimon) error
	UpdateStatus(ctx context.Context, d *Digimon) error
}

```

- `講解1 - struct`: 定義了數碼獸會有哪些的屬性，當程式裡面創建了數碼獸，就必定這些屬性都要擁有。
- `講解2 - repository interface`: 定義了 repository 層的各種方法，我們必須要按照這個定義實作，不然就會爆炸，呼叫的程式只能呼叫這些定義好的方法，不然也會爆炸 XD。
- `講解3 - usecase interface`: 與`講解-2`相同都是定義有哪些方法，不過他是定義 usecase 這個業務邏輯層。

### Repository 層 - 任何外部資料都我來管

定義好了 domain 層，就可以依照 domain 來設計 repository 層，

![](https://i.imgur.com/UtfWMqy.png)

可以看到雖然 diet 與 digimon 目前都是用 PostgreSQL 來實作，但是我們都將 repository 獨立拉出來，以便將不同的 DB 功能做區分。

以下為核心部分:

```golang
// ... 其他程式碼

// !!! 講解1 !!!
type postgresqlDigimonRepository struct {
	db *sql.DB
}

// !!! 講解2 && 講解3 !!!
// NewpostgresqlDigimonRepository ...
func NewpostgresqlDigimonRepository(db *sql.DB) domain.DigimonRepository {
	return &postgresqlDigimonRepository{db}
}

// !!! 講解4 !!!
func (p *postgresqlDigimonRepository) GetByID(ctx context.Context, id string) (*domain.Digimon, error) {
	row := p.db.QueryRow("SELECT id, name, status FROM digimons WHERE id =$1", id)
	d := &domain.Digimon{}
	if err := row.Scan(&d.ID, &d.Name, &d.Status); err != nil {
		logrus.Error(err)
		return nil, err
	}
	return d, nil
}

// ... 其他程式碼
```

- `講解1 - 定義好需要哪些依賴注入(DI)`: 這是 Clean Architecture 的核心之一，如[DAY06](https://github.com/superj80820/2020-ithelp-contest/blob/master/DAY06)所述，`將依賴的事物由外部注入，而不是寫死在裡頭`。
- `講解2 - 設計一個DI注入的function`: 有了`講解1`的定義，我們還需要一個注入的 function，將 db `確實`注入。為什麼要這樣呢？因為你是可以這樣產生實例的`postgresqlDigimonRepository{}`，可以發現其實沒有帶 db，但實際上還是可以運行。我們可以透過此 function 來避免這個情形。
- `講解3 - 透過domain裡定義的interface來約束回傳值`: `domain.DigimonRepository`這個 interface 定義了有哪些方法`postgresqlDigimonRepository`必須要實作，如[DAY06](https://github.com/superj80820/2020-ithelp-contest/blob/master/DAY06)所述，有了此 interface 才能讓呼叫的程式在`還沒run起來`就知道`哪些呼叫方法存在`。
- `講解4 - 實作domain.DigimonRepository interface定義的方法`: `GetByID`要確實符合 interface 定義的方法，不然 Golang 在運行前就會報錯。

### Usecase 層 - 業務邏輯的管轄處

![](https://i.imgur.com/BVDeh0X.png)

```go
// ... 其他程式碼

// !!! 講解1 !!!
type digimonUsecase struct {
	digimonRepo domain.DigimonRepository
}

// !!! 講解2 !!!
// NewDigimonUsecase ...
func NewDigimonUsecase(digimonRepo domain.DigimonRepository) domain.DigimonUsecase {
	return &digimonUsecase{
		digimonRepo: digimonRepo,
	}
}

// !!! 講解3 !!!
func (du *digimonUsecase) GetByID(ctx context.Context, id string) (*domain.Digimon, error) {
	aDigimon, err := du.digimonRepo.GetByID(ctx, id)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return aDigimon, nil
}

// ... 其他程式碼
```

- `講解1 - 定義依賴注入(DI)所需的Repository層`: 注入了之後我們再對 Repository 層做各種邏輯的操作，要注意的是，雖然目前只有注入所屬 PostgreSQL 的 repository，但如果 digimon 有很多不同的來源，比如說 MongoDB、Microservice 等等，我們可以注入更多 repository 來操作。
- `講解2 - 設計確實注入的function`: 與 repository 層一樣，需要一個 function 來要求要注入哪些 repository。
- `講解3 - 依照domain.DigimonRepository interface來實作`: 與 repository 層一樣，要以 interface 規範來實作，不過這裡是實作`業務邏輯`。

### Delivery 層 - 交付業務邏輯給引擎的跑腿工

![](https://i.imgur.com/AJ6zmsa.png)

看到這邊大家應該已經發現，Clean Architecture 的重點就是:

1. 定義好個層介面
2. 依賴注入注入再注入
3. 利用各種注入的實體來實作

而 delivery 層就是在注入 usecase 層的實體，

```go
// ... 其他程式碼

// !!! 講解1 !!!
// DigimonHandler ...
type DigimonHandler struct {
	DigimonUsecase domain.DigimonUsecase
	DietUsecase    domain.DietUsecase
}

// !!! 講解2 !!!
// NewDigimonHandler ...
func NewDigimonHandler(e *gin.Engine, digimonUsecase domain.DigimonUsecase, dietUsecase domain.DietUsecase) {
	handler := &DigimonHandler{
		DigimonUsecase: digimonUsecase,
		DietUsecase:    dietUsecase,
	}

	e.GET("/api/v1/digimons/:digimonID", handler.GetDigimonByDigimonID)
	e.POST("/api/v1/digimons", handler.PostToCreateDigimon)
	e.POST("/api/v1/digimons/:digimonID/foster", handler.PostToFosterDigimon)
}

// !!! 講解3 !!!
// PostToCreateDigimon ...
func (d *DigimonHandler) PostToCreateDigimon(c *gin.Context) {
    // !!! 講解3-1 !!!
	var body swagger.DigimonInfoRequest
	if err := c.BindJSON(&body); err != nil {
        logrus.Error(err)
        // !!! 講解3-2 !!!
		c.JSON(500, &swagger.ModelError{
			Code:    3000,
			Message: "Internal error. Parsing failed",
		})
		return
	}
	aDigimon := domain.Digimon{
		Name: body.Name,
	}
	if err := d.DigimonUsecase.Store(c, &aDigimon); err != nil {
        logrus.Error(err)
        // !!! 講解3-2 !!!
		c.JSON(500, &swagger.ModelError{
			Code:    3000,
			Message: "Internal error. Store failed",
		})
		return
	}

    // !!! 講解3-3 !!!
	c.JSON(200, swagger.DigimonInfo{
		Id:     aDigimon.ID,
		Name:   aDigimon.Name,
		Status: aDigimon.Status,
	})
}

// ... 其他程式碼
```

- `講解1 - 定義依賴注入(DI)所需的Usecase層`: 如 usecase 層，值得注意的是，由於 digimon 相關的 Restful API 除了有用到 DigimonUsecase，還有用到 DietUsecase，所以必須都注入進來，`並沒有規定只能注入digimon相關的usecase`。
- `講解2 - 將Server引擎丟進來並且設定`: 這裡比較特別，是透過把 Golang-Gin 的 Server 引擎傳進 function 內，再把各個 handler 與 route 交付(delivery)綁定。
- `講解3 - 透過`[swagger-generator](https://github.com/swagger-api/swagger-codegen)`來解析各個HTTP傳遞`: 這裡終於要用到[swagger-generator](https://github.com/swagger-api/swagger-codegen)的好處，可以看到不論是輸入處`講解3-1`、回傳錯誤處`講解3-2`、成功回傳處`講解3-3`都可以透過`swagger.定義好的介面`來解析。這樣就省去了很多定義時間，並且也減少了很多定義不小心寫錯的可能性！

## 最後，把一切透過 `cmd/main.go` 跑起來吧！

```go
// ... 其他程式碼

func main() {
	logrus.Info("HTTP server started")

    // !!! 講解1 !!!
	restfulHost := viper.GetString("RESTFUL_HOST")
	restfulPort := viper.GetString("RESTFUL_PORT")
	dbHost := viper.GetString("DB_HOST")
	dbDatabase := viper.GetString("DB_DATABASE")
	dbUser := viper.GetString("DB_USER")
	dbPassword := viper.GetString("DB_PASSWORD")

    // !!! 講解2 !!!
	db, err := sql.Open(
		"postgres",
		fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbUser, dbPassword, dbDatabase),
	)
	if err != nil {
		logrus.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		logrus.Fatal(err)
	}

	r := gin.Default()

    // !!! 講解3 !!!
	digimonRepo := _digmonRepo.NewpostgresqlDigimonRepository(db)
	dietRepo := _dietRepo.NewPostgresqlDietRepository(db)

    // !!! 講解4 !!!
	digimonUsecase := _digimonUsecase.NewDigimonUsecase(digimonRepo)
	dietUsecase := _dietUsecase.NewDietUsecase(dietRepo)

    // !!! 講解5 !!!
	_digimonHandlerHttpDelivery.NewDigimonHandler(r, digimonUsecase, dietUsecase)

    // !!! 講解6 !!!
	logrus.Fatal(r.Run(restfulHost + ":" + restfulPort))
}

// ... 其他程式碼
```

寫了那麼多依賴注入(DI)，最後我們需要再`main.go`真的將這些實體注入，

- `講解1 - 將config從外部依賴注入(DI)`: 我們統一由 main.go 來讀取 config，因為這個 config 是`外部`的，不該直接寫在各層裡頭，不然我們很難知道此層到底用了哪些 config，以避免：「嗯嗯？！怎麼爆炸了(查了 3 個小時後)，喔 damn 這層原來有要用這個 config 喔 Q 口 Q！」的狀況。
- `講解2 - 初始化DB`: 凡`外部`依賴的東西，都在 main.go 中完成，不該在各層中完成，以避免不知道誰已經初始化了誰還沒。
- `講解3 - 將外部依賴透過依賴注入(DI)來注入Repository層`: 也是因為`外部`的關係，我們也該將 DB 從外部注入，我們可以更可視此 DB 被哪些 repository 層使用了。除了 DB 以外其他 Microservice 的 caller、NoSQL 都該由外部注入，以便可視。
- `講解4 - 將Repository層透過依賴注入(DI)來注入Usecase層`: 同`講解3`，對 usecase 層來說 repository 層也是外部依賴，注入下去！
- `講解5 - 將Usecase層透過依賴注入(DI)來注入Delivery層`: 這裡要注意，除了注入以外，還有將 Golang-Gin 的引擎丟入，讓 delivery 層來綁定。
- `講解 6 - 將Golang-Gin跑起來！`

進入到我們 DAY07 資料夾底下，使用 docker-compose 把 Golang-Server 與 DB 跑起來吧！

```bash
$ cd DAY07
$ docker-compose up
```

![](https://i.imgur.com/s42lh1p.png)

Work!

## 透過[Insomnia Designer](https://insomnia.rest/products/designer/)測試一下

創建數碼蛋，

![](https://i.imgur.com/6mrhFsX.png)

查看數碼蛋狀態，

![](https://i.imgur.com/T0osPBU.png)

培育數碼獸，

![](https://i.imgur.com/DWNakcv.png)

如果你有安裝[Postico](https://eggerapps.at/postico/)，可以去看看 diets table 是不是真的有吃了食物，

![](https://i.imgur.com/jdfwHFc.png)

嗯！滿滿的蘋果，看我還不飽炸你，亞古獸 XD ～

![](https://i.imgur.com/ZKDEbdn.png)

## 參考

- [bxcodec/go-clean-arch](https://github.com/bxcodec/go-clean-arch)
- [Trying Clean Architecture on Golang](https://medium.com/hackernoon/golang-clean-archithecture-efd6d7c43047)
