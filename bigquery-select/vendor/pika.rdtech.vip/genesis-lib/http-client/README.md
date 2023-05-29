http-client
===

## Core
reference : https://github.com/syhlion/greq

## Method
* Get(ctx context.Context, uri string, param map[string]string) (r *Response)
* Post(ctx context.Context, uri string, body []byte) (r *Response)
* Put(ctx context.Context, uri string, body []byte) (r *Response)
* Delete(ctx context.Context, uri string, body []byte) (r *Response)
* GetWithHeader(ctx context.Context, uri string, param, header, cookie map[string]string) (r *Response)
* PostWithHeader(ctx context.Context, uri string, body []byte, header, cookie map[string]string) (r *Response)
* PutWithHeader(ctx context.Context, uri string, body []byte, header, cookie map[string]string) (r *Response)
* DeleteWithHeader(ctx context.Context, uri string, body []byte, header, cookie map[string]string) (r *Response)
* GetTimeOut() (timeout time.Duration)
* GetHeader(key string) (value string)
* GetHost() string
* GetScheme() string
* GetAddr() (addr string)
* GetWithBody(ctx context.Context, uri string, param map[string]string, body []byte, header, cookie map[string]string) (resp Response)

## Options
* SetAddr(addr string)
* SetSSL(ssl string)
* SetHost(host string)
* SetTimeOut(timeout time.Duration)
* SetHeader(key string, value string)
* SetDebug(open bool)

## Quick Start
```go
package main

import (
	"context"
	"fmt"

	httpClient "pika.rdtech.vip/genesis-lib/http-client"
)

func main () {
    //new client with number of workers
	conn := httpClient.New(context.Background(), 10)

	//GET
	getRes := conn.Get(context.Background(), "https://github.com/syhlion/greq", nil)
	fmt.Println(getRes.Code, getRes.Body.String(), getRes.Err)
}
```

## Options Usage

```go
package main

import (
	"context"
	"fmt"
	"time"

	httpClient "pika.rdtech.vip/genesis-lib/http-client"
)

func main () {
    //new client with options
    conn := httpClient.New(
        context.Background(),
        30,
        httpClient.SetSSL("https"),
        httpClient.SetHost("github.com"),
        httpClient.SetHeader(map[string]string{"testheader": "TH"}),
        httpClient.SetTimeOut(10*time.Second),
    )

    //GET
    getRes := conn.Get(context.Background(), "/syhlion/greq", nil)
    fmt.Println(getRes.Code, getRes.Body.String(), getRes.Err)
}
```

## Trace Logs

```go
package main

import (
	"context"
	"fmt"

	httpClient "pika.rdtech.vip/genesis-lib/http-client"
)

func main () {
    //new client with number of workers
    conn := httpClient.New(context.Background(), 10)

    //GET
    getRes := conn.Get(context.Background(), "https://github.com/syhlion/greq", nil)

    //print http trace
    getRes.Trace.Print("YOUR TAG")

    //print yourself
    fmt.Println("[Method]", getRes.Trace.Method)
    fmt.Println("[Total Time]", getRes.Trace.Total)
}
```