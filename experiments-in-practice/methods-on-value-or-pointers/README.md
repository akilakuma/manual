methods-on-value-or-points
===

## method 是基於value或者points探討

``` go
func (u User) GetName() string {}
func (u *User) GetName() string {}
```

## 結論
1. 不論傳value or pointers，呼叫的method就那一個，呼叫幾次，method的address都不變，表示不會另外增生method。


2. 不論傳value or pointers，在method裡面，拿到的struct對象都是新的struct位址，但是傳value的話，裡面的field都會是新的。而傳pointers的話，雖然struct是新的，但field內容都指到原本的地方。

```
原本預期0xc000000e030會有3個，並非0xc000000e030、0xc000000e038、0xc000000e040
```

## 執行結果

![](https://i.imgur.com/cK6jJC9.png)

### update time
2020-05-21
