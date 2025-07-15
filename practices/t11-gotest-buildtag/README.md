利用build tag協助go test
===

有兩個名稱一樣的method：GetInfo()，分別放在a.go和a_plus.go兩支檔案。

利用指定build tag的方式，讓go test執行的到我們指定的method。

```
go test -v

# 結果是 plus
```

```
go test -v -tags debug

# 結果是 origin
```
