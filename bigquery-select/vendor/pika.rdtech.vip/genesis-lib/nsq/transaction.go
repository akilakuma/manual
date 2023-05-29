package nsq

import (
	"github.com/nsqio/go-nsq"
)

//ProducerTransaction is returned by the async publish methods
type ProducerTransaction struct {
	t chan *nsq.ProducerTransaction
}

//NewProducerTransaction NewProducerTransaction
func NewProducerTransaction() *ProducerTransaction {
	return &ProducerTransaction{t: make(chan *nsq.ProducerTransaction)}
}
