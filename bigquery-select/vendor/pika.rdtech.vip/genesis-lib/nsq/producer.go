package nsq

import (
	"errors"
	"sync"
	"time"

	json "pika.rdtech.vip/genesis-lib/json"
	nsqio "github.com/nsqio/go-nsq"
)

//Producer is a high-level type to publish to NSQ.
type Producer interface {
	Ping() error
	Address() string
	Stop()
	Publish(topic string, body []byte) error
	PublishJSON(topic string, body interface{}) error
	MultiPublish(topic string, body [][]byte) error
	MultiPublishJSON(topic string, body []interface{}) error
	PublishAsync(topic string, body []byte, done *ProducerTransaction, args ...interface{}) error
	PublishAsyncJSON(topic string, body interface{}, done *ProducerTransaction, args ...interface{}) error
	MultiPublishAsync(topic string, body [][]byte, done *ProducerTransaction, args ...interface{}) error
	MultiPublishAsyncJSON(topic string, body []interface{}, done *ProducerTransaction, args ...interface{}) error
	DeferredPublish(topic string, delay time.Duration, body []byte) error
	DeferredPublishJSON(topic string, delay time.Duration, body interface{}) error
	DeferredPublishAsync(topic string, delay time.Duration, body []byte, done *ProducerTransaction, args ...interface{}) error
	DeferredPublishAsyncJSON(topic string, delay time.Duration, body interface{}, done *ProducerTransaction, args ...interface{}) error
}

type producer struct {
	p          *nsqio.Producer
	topics     *sync.Map
	topicCheck bool
}

//Ping() ping u know
func (w *producer) Ping() error {
	return w.p.Ping()
}

// Address returns the address of the Producer
func (w *producer) Address() string {
	return w.p.String()
}

// Stop initiates a graceful stop of the Producer (permanent)
func (w *producer) Stop() {
	w.p.Stop()
}

// Publish synchronously publishes a message body to the specified topic, returning
func (w *producer) Publish(topic string, body []byte) error {
	if err := w.checkTopics(topic); err != nil {
		return err
	}

	return w.p.Publish(topic, body)
}

// PublishJSON PublishJSON
func (w *producer) PublishJSON(topic string, body interface{}) error {
	if err := w.checkTopics(topic); err != nil {
		return err
	}

	b, err := json.Marshal(body)
	if err != nil {
		return err
	}
	return w.p.Publish(topic, b)
}

// MultiPublish synchronously publishes a slice of message bodies to the specified topic, returning
func (w *producer) MultiPublish(topic string, body [][]byte) error {
	if err := w.checkTopics(topic); err != nil {
		return err
	}

	return w.p.MultiPublish(topic, body)
}

// MultiPublishJSON synchronously publishes a slice of message bodies to the specified topic, returning
func (w *producer) MultiPublishJSON(topic string, body []interface{}) error {
	if err := w.checkTopics(topic); err != nil {
		return err
	}

	var b [][]byte
	for range body {
		tmp, err := json.Marshal(body)
		if err != nil {
			return err
		}
		b = append(b, tmp)
	}

	return w.p.MultiPublish(topic, b)
}

// PublishAsync publishes a message body to the specified topic
func (w *producer) PublishAsync(topic string, body []byte, done *ProducerTransaction, args ...interface{}) error {
	if err := w.checkTopics(topic); err != nil {
		return err
	}

	return w.p.PublishAsync(topic, body, done.t, args...)
}

//PublishAsyncJSON PublishAsyncJSON
func (w *producer) PublishAsyncJSON(topic string, body interface{}, done *ProducerTransaction, args ...interface{}) error {
	if err := w.checkTopics(topic); err != nil {
		return err
	}

	b, err := json.Marshal(body)
	if err != nil {
		return err
	}
	return w.p.PublishAsync(topic, b, done.t, args...)
}

// MultiPublishAsync publishes a slice of message bodies to the specified topic
func (w *producer) MultiPublishAsync(topic string, body [][]byte, done *ProducerTransaction, args ...interface{}) error {
	if err := w.checkTopics(topic); err != nil {
		return err
	}

	return w.p.MultiPublishAsync(topic, body, done.t, args...)
}

// MultiPublishAsyncJSON MultiPublishAsyncJSON
func (w *producer) MultiPublishAsyncJSON(topic string, body []interface{}, done *ProducerTransaction, args ...interface{}) error {
	if err := w.checkTopics(topic); err != nil {
		return err
	}

	var b [][]byte
	for range body {
		tmp, err := json.Marshal(body)
		if err != nil {
			return err
		}
		b = append(b, tmp)
	}

	return w.p.MultiPublishAsync(topic, b, done.t, args...)
}

// DeferredPublish synchronously publishes a message body to the specified topic
func (w *producer) DeferredPublish(topic string, delay time.Duration, body []byte) error {
	if err := w.checkTopics(topic); err != nil {
		return err
	}

	return w.p.DeferredPublish(topic, delay, body)
}

// DeferredPublish DeferredPublishJSON
func (w *producer) DeferredPublishJSON(topic string, delay time.Duration, body interface{}) error {
	if err := w.checkTopics(topic); err != nil {
		return err
	}

	b, err := json.Marshal(body)
	if err != nil {
		return err
	}
	return w.p.DeferredPublish(topic, delay, b)
}

// DeferredPublishAsync publishes a message body to the specified topic
func (w *producer) DeferredPublishAsync(topic string, delay time.Duration, body []byte,
	done *ProducerTransaction, args ...interface{}) error {
	if err := w.checkTopics(topic); err != nil {
		return err
	}

	return w.p.DeferredPublishAsync(topic, delay, body, done.t, args...)
}

// DeferredPublishAsync publishes a message body to the specified topic
func (w *producer) DeferredPublishAsyncJSON(topic string, delay time.Duration, body interface{},
	done *ProducerTransaction, args ...interface{}) error {
	if err := w.checkTopics(topic); err != nil {
		return err
	}

	b, err := json.Marshal(body)
	if err != nil {
		return err
	}
	return w.p.DeferredPublishAsync(topic, delay, b, done.t, args...)
}

// Check topic was registed
func (w *producer) checkTopics(topic string) error {
	if w.topicCheck {
		if _, ok := w.topics.Load(topic); !ok {
			return errors.New(`Topic "` + topic + `" not registed`)
		}
	}

	return nil
}
