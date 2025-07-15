package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gomodule/redigo/redis"
)

func main() {
	// 打開redis
	// redisConnect, err := redis.Dial("tcp", "127.0.0.1:6379")
	// if err != nil {
	// 	fmt.Println("somthing error")
	// }
	// defer redisConnect.Close()

	redisPool := &redis.Pool{
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

	redisPool.Get()

	go testWatch(redisPool.Get())
	go testWatch(redisPool.Get())

	log.Println("done ")
	select {}
}

func testWatch(redisC redis.Conn) {
	for i := 0; i < 100000; i++ {
		log.Println(i)
		_, err1 := redisC.Do("Watch", "user_balance", 12345)
		if err1 != nil {
			log.Println(err1.Error())
		}
		_, err2 := redis.String(redisC.Do("HGET", "user_balance", 12345))
		if err2 != nil {
			fmt.Println(err2.Error())
		}

		errM := redisC.Send("MULTI")

		if errM != nil {
			fmt.Println(errM.Error())
		}

		errT1 := redisC.Send("HSET", "user_balance", 12345, i)
		if errT1 != nil {
			fmt.Println(errT1.Error())
		}

		_, errE := redisC.Do("EXEC")
		if errE != nil {
			fmt.Println(errE.Error())
		}
		time.Sleep(10 * time.Microsecond)
	}
}
