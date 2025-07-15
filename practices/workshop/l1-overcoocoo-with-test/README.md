第一個作業，練習寫寫測試
===

## 指令

### 詳細測試結果
``` bash
go test golang-advance-practice/workshop/l1-overcoocoo-with-test/restaurant -v
```

### 輸出覆蓋率檔案
``` bash
go test golang-advance-practice/workshop/l1-overcoocoo-with-test/restaurant -coverprofile=c.out && go tool cover -html=c.out
```
