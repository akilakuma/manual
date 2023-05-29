# inlog



## Install

`go get -u pika.rdtech.vip/eden-lib/inlog`



## Log Level

1. Default(0) 日誌條目沒有指定的嚴重性級別
2. Debug(100) 調試或跟踪信息
3. Info(200) 常規信息，例如正在進行的狀態或性能
4. Notice(300) 正常但重要的事件，例如啟動，關閉或配置更改
5. Warn(400) 警告事件可能會導致問題
6. Error(500) 錯誤事件可能會導致問題
7. Critical(600) 嚴重事件會導致更嚴重的問題或中斷
8. Alert(700) 一個人必須立即採取行動
9. Emergency(800) 一個或多個系統無法使用


## Usage

``` golang
package main

var (
    //注意 此 log 可以重複使用，不需要重複建構
    log = inlog.New()

    //inerror 也可以這樣使用 不用每次都需要帶入 service name
    inerr = inerror.New(inerror.AppleTree)
)

func main(){


    //一般用法
    log.Debug("yao 46")
    log.Debugf("yao %d",46)
    //output: {"level":"DEBUG","message":"yao 46","severity":100,"time":"2018-11-28T14:42:57.496067+08:00"}



    var err error
    //使用 inerror 裡面會做特別處理 
    err = inerr.Wrap(777,"yaoming 請客",nil,nil)

    log.WithError(err).Notice("到屋馬集合")
    //output: {"code":120010777,"err_msg":"yaomin 請客","extrainfo":null,"level":"NOTICE","message":"到屋馬集合","origin_err":"","service":"APPLETREE","severity":300,"time":"2018-11-28T14:47:53.417885+08:00"}

    
    //不使用 inerror 會如下(一般常規的 error)
    err = errors.New("yaoming 又要請客")

    log.WithError(err).Alert("到老乾杯集合")
    //output: {"error":"yaomin 又要請客","level":"ALERT","message":"到老乾杯集合","severity":700,"time":"2018-11-28T14:47:53.41797+08:00"}


    //使用 grpc 接口的 error(假設 grpc server 都有照 inerror格式回傳,parser 不出來 會直接 formate 9999處理)
    ret,err:=grpcClient.Order(context.Backgrounc(),&Spec{
        Msg:"Hello",
    })
    if err != nil {
        log.WithGrpcError(err).Alert("grpc error")
    }
    //output: {"code":120010777,"err_msg":"yaomin 請客","extrainfo":null,"level":"NOTICE","message":"grpc error","origin_err":"","service":"APPLETREE","severity":300,"time":"2018-11-28T14:47:53.417885+08:00"}

}

```
