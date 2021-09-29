
package main

import (
"fmt"
"strconv"

"github.com/go-redis/redis"
)

const (
	addr           = "127.0.0.1:6379"
	ConstUserPoint = "DUCK_user_point"
)

var luaScript = `

		local last_point = 0
		local point = tonumber(ARGV[1])
		local sum = 0
		local isSuccess = false

		last_point = redis.call('HGET', KEYS[1], KEYS[2])

		if last_point == false then
			if point < 0 then 
				return  {tostring(isSuccess), tostring(0)}
			else
				isSuccess = true
				redis.call('HSET', KEYS[1], KEYS[2], point)
				return  {tostring(isSuccess), tostring(point)}
			end
		end


		sum  = last_point + point
		if  sum > 0 then 
			isSuccess = true
			redis.call('HSET', KEYS[1], KEYS[2], sum)
		else
			isSuccess = false
			-- 回傳目前餘額
			sum = last_point
		end

		return  {tostring(isSuccess), tostring(sum)}
	`

var luaScript1 *redis.Script

func main() {
	appleTreeRClient, err := newRedisConnection(addr, 10, 10, 0)
	if err != nil {
		fmt.Println("redis connect err:", err)
		return
	}

	b := RedisGetUserPoint(appleTreeRClient, 23456, "abcdef")
	fmt.Println("balance is :", b)

	RedisAddUserPoint(appleTreeRClient, 1, 23456, "abcdef")
}

// 取得對應的redis key
// user_id + hash_id
// hash_id 等同於一個遊戲連線使用，因此也許是同一款遊戲但多開
func getUserKey(userId int, hashId string) string {
	return strconv.Itoa(userId) + "_" + hashId
}

// RedisGetUserPoint 取得使用者目前分數
func RedisGetUserPoint(conn *redis.Client, userId int, hashId string) int {

	result := conn.HGet(ConstUserPoint, getUserKey(userId, hashId))
	userPoint, getErr := result.Int()
	if getErr != nil {
		fmt.Println("getErr:", getErr)
	}

	return userPoint
}

// RedisAddUserPoint 加減使用者目前分數
// bool:回傳操作是否成功，int：回傳目前餘額
func RedisAddUserPoint(conn *redis.Client, point, userId int, hashId string) {

	scriptHash, err1 := conn.ScriptLoad(luaScript).Result()

	result := conn.EvalSha(scriptHash, []string{ConstUserPoint, getUserKey(userId, hashId)}, point)

	r, getResultErr := result.Result()

	fmt.Println(r, getResultErr)

	a := r.([]interface{})
	r1, _ := strconv.ParseBool(a[0].(string))
	r2, _ := strconv.ParseFloat(a[1].(string), 64)
	r3, _ := strconv.ParseFloat(a[2].(string), 64)

	fmt.Println(err1, r, getResultErr, r1, r2, r3)
}


func newRedisConnection(addr string, maxIdle int, maxConn int, db int) (client *redis.Client, err error) {
	client = redis.NewClient(&redis.Options{
		Addr:       addr,
		DB:         db,
		MaxRetries: 5,
	})

	_, err = client.Ping().Result()
	if err != nil {
		ex := make(map[string]interface{})
		ex["addr"] = addr
		ex["db"] = db
		fmt.Println("connect err:", err)
		return nil, err
	}

	return
}
