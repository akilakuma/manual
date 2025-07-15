interface with golang
===


結論：
1. 假如有在interface裡面宣告擁有A、B、C共3個函式，實際上沒有全寫出使用interface的 func Ａ、func B、func的話，
在編譯時不會報錯，不過間接也等於這個interface廢掉了，因為想使用的時候，傳入參數一定報錯，編譯不會過。

>>> interface宣告幾個function名稱，就要寫出幾個function出來(實作)，不論傳入的stuct長得怎麼樣。
>>> 加一個新的struct，若該struct想使用某個介面，就必須寫出介面的所有function出來。

2. 我原先不了解的地方，func (c circle) area() 這樣的寫法，就目前看來，就好像表示circle 這個struct有一個叫做area的method

>>> var c circle
    c.area()

>>> circle有個area的method，而且這個area只專屬於circle，area內可以使用在circle裡面宣告的變數們
>>> 所以有這樣的綁定關係，一但修改了struct，與struct綁定的method有可能受到影響而需要調整。



OK，上述兩點跟interface其實無太大相關，這個東西關鍵字叫做method in struct
