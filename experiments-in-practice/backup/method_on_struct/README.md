function in struct with golang
===


##說明：
藉由gorutine 去呼叫test2()兩次，test2 有呼叫 function in struct 的部分。
第一次在print之前先進行sleep，看會不會被第二次進來的test2影響到。

##執行結果：

YOOOO
0xc420096018
LU!
0xc4200ae000
NANANA
0xc42000c010


##結論：
使用fucntion test2 裡面的Foo是指向ㄧ個新的記憶體，所以不會互相影響yo!
