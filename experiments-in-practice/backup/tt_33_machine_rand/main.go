package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/garyburd/redigo/redis"
)

// Total 亂數產生數量
const Total = 1000000

// Cut 切割的數量
const Cut = 50000

// RandNum 隨機亂數產生
func RandNum() int {

	var out bytes.Buffer
	cmd := exec.Command("head", "-c", "20", "/dev/hwrng")
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		//本機mac開發無/dev/hwrng
		cmd := exec.Command("head", "-c", "20", "/dev/urandom")
		cmd.Stdout = &out
		err2 := cmd.Run()
		if err2 != nil {
			log.Print("error", "error_test: "+err.Error())
		}
	}
	randNum := binary.BigEndian.Uint32(out.Bytes())
	return int(randNum)
}

func connectRedis() redis.Conn {
	redisPool := &redis.Pool{
		Wait:        true,
		MaxIdle:     10,
		MaxActive:   1000,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "redis:6379")
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

	return redisPool.Get()

}

func main() {

	r := connectRedis()

	var (
		temp    []string
		content string
	)
	for i := 0; i < Total; i++ {
		a := RandNum()
		// fmt.Print(a, " ")
		temp = append(temp, strconv.Itoa(a))

		if i%Cut == 0 {

			errM := r.Send("MULTI")
			if errM != nil {
				fmt.Println(errM.Error())
			}

			content = strings.Join(temp[:], ",")

			errT1 := r.Send("SADD", i, content)
			if errT1 != nil {
				fmt.Println(errT1.Error())
			}

			_, errT2 := r.Do("EXEC")
			if errT2 != nil {
				fmt.Println(errT2.Error())
			}

			temp = []string{}
		}
	}
}
