# inerror


## Install

`go get pika.rdtech.vip/eden-lib/inerror`


## Api


* inerror.Error 有實作error interface


* Wrap(errorCode int,errorMsg string,errorService inerror.Service,errorInfo map[string]interface{},orignErr error) inerror.Error

`errorInfo 為pitaya的 extrainfo，為選填，如不需要可帶入nil`

`originErr 為某些情境，不想把原始error覆蓋掉，但是又不想暴露出去，把原始error帶入這裡，如不需要可帶入nil`


* Parse(errorString string) inerror.Error

`帶入由 Wrap 包出來的error tostring後，可再反解回inerror.Error struct`


## Usage

grpc server

```
ers:=inerror.New(inerror.AppleTree)

func (t *Test) Test(ctx context.Context, er *pb.Request) (p *pb.Response, err error) {
    err:=db.Scan(...)
    if err != nil {
        return nil, ers.Wrap(1,"db query error",nil,err) 
    }
    return &pb.Response{
        Success: &pb.Success{
            Test: "hihi",
        },
    }, nil
}


```


grpc gateway

```
func CustomHTTPError(ctx context.Context, _ *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, _ *http.Request, err error) {
    const fallback = `{"error": "failed to marshal error message"}`

    w.Header().Set("Content-type", marshaler.ContentType())
    w.WriteHeader(runtime.HTTPStatusFromCode(grpc.Code(err)))

    fmt.Printf("%#v", err)

    var ebdy errorBody
    er := grpc.ErrorDesc(err)
    realErr:=inerror.Parse(er)
    jErr := json.NewEncoder(w).Encode(realErr)

    if jErr != nil {
        w.Write([]byte(fallback))
    }
}
```
