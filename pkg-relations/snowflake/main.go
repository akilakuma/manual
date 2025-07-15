package main

import (
	"time"

	"log"

	"github.com/bwmarrin/snowflake"
)

/*
	另外一種選擇是sonyflake
	https://github.com/sony/sonyflake
*/

func main() {
	var snowflakeNode *snowflake.Node
	snowflakeNode = CreateSnowflake(1)

	for {
		cartId := snowflakeNode.Generate().Int64()

		log.Println(cartId)
		time.Sleep(1 * time.Second)

	}
}

// CreateSnowflake 建立snowflake
func CreateSnowflake(nodeId int64) *snowflake.Node {
	snowflake, err := snowflake.NewNode(nodeId)

	if err != nil {
		panic(err)
	}

	log.Println("Snowflake 建立成功")

	return snowflake
}
