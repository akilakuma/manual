package main

import (
	"context"
	"encoding/json"
	"fmt"

	"cloud.google.com/go/pubsub"
)

func main() {
	bqPubSub()
}

// bqPubSub Google pub sub 示範method
func bqPubSub() {
	fmt.Println("start pub sub")
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, "rd2-test")
	if err != nil {
		// TODO: Handle error.
		fmt.Println("bqPubSub err: ", err)
		return
	}

	topic := client.Topic("dev_bq_v1")

	for i := 0; i < 10; i++ {
		d, _ := json.Marshal(&BqMessage{
			UserId:      1552463678243115008,
			UserName:    "dev_bobcat_set",
			UserLabel:   "vip",
			LabelWeight: 100,
			CreatedAt:   1658937600,
			Serial:      2 + i,
		})

		_, pubErr := topic.Publish(ctx, &pubsub.Message{Data: d}).Get(ctx)
		if pubErr != nil {
			fmt.Println("puma publish:", pubErr.Error())
		}
	}
	fmt.Println("pub done")
}

type BqMessage struct {
	UserId      int64  `json:"user_id"`
	UserName    string `json:"user_name"`
	UserLabel   string `json:"user_label"`
	LabelWeight int    `json:"label_weight"`
	CreatedAt   int64  `json:"created_at"`
	Serial      int    `json:"serial_num"`
}

//{
//"user_id":      1552463678243115008,
//"user_name":    "dev_bobcat_set",
//"user_label":   "vip",
//"label_weight": 100,
//"created_at":   1658937600,
//"serial_num":      1,
//}
