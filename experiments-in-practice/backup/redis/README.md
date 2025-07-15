redis with golang
===

### os
- 在mac

### 前置環境
- go
- docker

### package
連結
[github.com/gomodule/redigo/redis](https://github.com/gomodule/redigo)

安裝指令
    go get github.com/gomodule/redigo/redis

### redis 環境(連本機)

#### 1. redis 服務
- 使用laradock
- 執行指令

        docker-compose up -d redis

#### 2. phpRedisAdmin
- dockerhub 上有人做好的可以用！
- 執行指令

        docker run --rm -it -e REDIS_1_HOST=192.168.152.84 -e REDIS_1_NAME=MyRedis -p 80:80 erikdubbelboer/phpredisadmin

**192.168.152.2 是本機IPv4 位置**

![](https://i.imgur.com/Tt2PkrB.png =600x150)

```
```

```
```

**成功執行結果**

![](https://i.imgur.com/K16DXPc.png =600x300)

