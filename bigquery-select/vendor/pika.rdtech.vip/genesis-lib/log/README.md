# log

### Log Level
* Default(0) 日誌條目沒有指定的嚴重性級別
* Debug(100) 調試或跟蹤信息
* Info(200) 常規信息，例如正在進行的狀態或性能
* Notice(300) 正常但重要的事件，例如啟動，關閉或配置更改
* Warn(400) 警告事件可能會導致問題
* Error(500) 錯誤事件可能會導致問題
* Critical(600) 嚴重事件會導致更嚴重的問題或中斷
* Alert(700) 一個人必須立即採取行動
* Emergency(800) 一個或多個系統無法使用

### Use Export log

```golang
package main

import (
    "pika.rdtech.vip/genesis-lib/log"
)

func main() {
    log.SetLevel(log.DebugLevel)
    log.SetFormatter(&log.JSONFormatter{})

    log.Warn("ABC")
    log.Error("EDF")

    ll := l.WithFields(log.Fields{
        "omg":    true,
        "number": 122,
    })

    ll.Debug("TEST")

    ll.WithFields(log.Fields{
        "animal": "walrus",
        "size":   10,
    }).Info("Info")
}
```

### Use New log

```golang
package main

import (
    "pika.rdtech.vip/genesis-lib/log"
)

func main() {
    l := log.New()

    l.SetLevel(log.DebugLevel)
    l.SetFormatter(&log.JSONFormatter{})

    l.Warn("ABC")
    l.Error("EDF")

    ll := l.WithFields(log.Fields{
        "omg":    true,
        "number": 122,
    })

    ll.Debug("TEST")

    ll.WithFields(log.Fields{
        "animal": "walrus",
        "size":   10,
    }).Info("Info")
}
```


### Use For InError

```golang
package main

import (
    "errors"
    "strconv"

    "pika.rdtech.vip/eden-lib/inerror"
    "pika.rdtech.vip/genesis-lib/log"
)

func main() {
    Log := log.New()
    Log.SetLevel(log.DebugLevel)

    ff := &FuckFormat{
        JsonF: &log.JSONFormatter{},
    }

    Log.SetFormatter(ff)

    ers := inerror.New(inerror.Sage)

    Log.WithError(ers.Wrap(9999, "ErrMsgSystem", map[string]interface{}{"test": "test"}, errors.New("origin_test"))).
        Error("test error")
}

type FuckFormat struct {
    JsonF *log.JSONFormatter
}

func (f *FuckFormat) Format(entry *log.Entry) ([]byte, error) {
    err, ok := entry.Data[log.ErrorKey]

    if !ok {
        return f.JsonF.Format(entry)
    }

    if inerr, ok := err.(inerror.Error); ok {
        entry.Data["code"] = inerr.Code
        entry.Data["err_msg"] = inerr.Msg
        entry.Data["extrainfo"] = ConvertInt64ToString(inerr.ExtraInfo)
        entry.Data["service"] = inerr.Service
        entry.Data["origin_err"] = inerr.OriginErr
        delete(entry.Data, log.ErrorKey)
    }

    return f.JsonF.Format(entry)
}

func ConvertInt64ToString(m map[string]interface{}) (result map[string]interface{}) {
    for key, value := range m {
        switch v := value.(type) {
        case int64:
            int64s := strconv.FormatInt(v, 10)
            m[key] = int64s
        default:
            continue
        }
    }

    return m
}

```