mysql with golang
===

### os
- 在mac

### 前置環境
- go
- docker

### package
連結
[github.com/jinzhu/gorm](https://github.com/jinzhu/gorm)

安裝指令
    go get github.com/jinzhu/gorm

### mysql 環境(連本機)

#### 1. mysql 服務
- 使用資料夾內的docker-compose.yml
- 執行指令

        docker-compose up -d local-mysql
        
#### 2. 使用 Sequel Pro
- 因為我受不了網頁版的DB managent
- https://www.sequelpro.com/
- 帳號、密碼、設定我寫在docker-compose.yml
帳號是root
密碼qwe123
預設DB是Test


### 備註
- laradock的mysql是v8，有點問題，下docker-compose，發現其實沒有跑起來
- docker hub 上有個mysqlserver，也很奇怪bind-address竟然設*，我不會設也跑不起來
- 所以只好自己做compose.yml


