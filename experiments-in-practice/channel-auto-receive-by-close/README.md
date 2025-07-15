channel-auto-receive-by-close
===

## 如果close channel，channel的接收端會不會自動拿到！ok


## 結論
會，因為要關閉goroutine，不需要額外往channel塞什麼資料作通知。


## 執行結果

![](https://i.imgur.com/IKka108.png)


### update time
2020-05-26
