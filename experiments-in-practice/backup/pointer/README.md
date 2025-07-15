pointer with golang
===

- golang 沒有call by refference這回事，像我這種剛從python、php、javascript轉過來的，可能特別容易踏進去這個坑。
[參考文章](https://medium.com/manjeaneer/go%E4%B8%8D%E6%98%AFpython-call-by-object-reference-980882711278)

- 以前學C和C++對指標有很大誤解，老師對不起，但你講的真的不親民。
```go
b := 10
var a *int = &b
```
和
```go
b := 10
a:= &b
```
和
```go
func main() {
    b := 10
    setFunc(&b)
}

func setFunc(a *int) {

}

```
- 以上三個都在說同樣的一件事，但我不知道為什麼以前透過function傳接球，腦袋轉不過來。

```go
var a *int = &b
```
- 將『指標(命名為a) 指向b』用法就這樣，沒有別的，沒事不會亂用或單獨用『&』『*』

- 『int』 跟 『*int』雖然長得87分像，但兩個完全不一樣， 『int』叫做整數，『*int』如上的用途的時候叫做指標，指的對像是整數。如果『*int』可以改叫做『pointerToInt』之類的關鍵字，我大一也不會被當掉了，大概吧。

- 改值是另一件事， *a = 10，這時候的『星號』已經暗指a在之前已經被宣告成指標的type，這點很重要，『星號』在這時候的用途英文叫做derefference，中文有人叫做『提取』，不管你是一層指標，還是像Ｃ語言有人有病寫了三層指標，都是存這個值得，一但用了『提取』(*a = 10)，不管在哪裡都會被改到。
