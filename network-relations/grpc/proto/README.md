## 產生gofile

cd 到account
```shell
protoc --proto_path=.. \
       --go_out=.. \
       --go_opt=paths=source_relative \
       --go-grpc_out=.. \
       --go-grpc_opt=paths=source_relative \
       account/account.proto
```

cd 到room
```shell
protoc --proto_path=.. \
       --go_out=.. \
       --go_opt=paths=source_relative \
       --go-grpc_out=.. \
       --go-grpc_opt=paths=source_relative \
       room/room.proto
```