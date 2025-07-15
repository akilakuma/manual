package main

import (
	"fmt"
	"log"
	Random "math/rand"
	"os"
	"strconv"
	"time"
	// "reflect"
	"github.com/gomodule/redigo/redis"
)

var (
	times = 10000000
)

func main() {

	fmt.Println("practice start ......")
	exitChan := make(chan struct{}, 0)

	// 打開redis
	conn := getConnectionPool()
	redisConnect := conn.Get()
	defer redisConnect.Close()

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "push":
			PushRedis(redisConnect, exitChan)
		case "pop":
			PopRedis(redisConnect)
		}
	}

	fmt.Println("practice end ......")
	<-exitChan
}

func PushRedis(redisC redis.Conn, exitChan chan struct{}) {
	s1 := Random.NewSource(time.Now().UnixNano())

	r1 := Random.New(s1)

	randTime := r1.Intn(10)

	// 結果寫回array
	for i := 0; i < times; i++ {
		time.Sleep(time.Duration(randTime) * time.Microsecond)
		_, err := redisC.Do("RPUSH", "data_queue", "num_"+strconv.Itoa(i))

		if err != nil {
			log.Println(err.Error())
		}
	}
	close(exitChan)
}

func PopRedis(redisC redis.Conn) {
	s1 := Random.NewSource(time.Now().UnixNano())

	r1 := Random.New(s1)

	randTime := r1.Intn(10)

	var (
		before string
		now    string
	)

	// 結果寫回array
	for i := 0; i < times; i++ {
		time.Sleep(time.Duration(randTime) * time.Microsecond)
		s, err := redis.String(redisC.Do("LPOP", "data_queue"))

		if err != nil {
			log.Println(err.Error())
		}

		now = s
		if before == now {
			log.Println("有出現pop出同樣的值")
			break
		}

		before = now
	}
}

// getConnectionPool 實質建立連線
func getConnectionPool() *redis.Pool {

	return &redis.Pool{
		Wait:        true,
		MaxIdle:     10,
		MaxActive:   1000,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "127.0.0.1:6379")
			if err != nil {
				log.Println(err.Error())
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			if err != nil {
				log.Println(err.Error())
			}
			return err
		},
	}
}
