NSQ
===

## Core

reference: https://github.com/nsqio/go-nsq

## Usage

#### Usage of Topic check

```go
import (
	"pika.rdtech.vip/genesis-lib/nsq"
)

func init() {
    config := nsq.NewConfig()
    config.SetDialTimeout(10 * time.Second)
    config.SetMaxAttempts(30)
    config.SetMaxRequeueDelay(500 * time.Millisecond)

    //[*optional] this setting will open auto connection check & repair consumer
    if err := config.SetConnCheck("127.0.0.1:4147", time.Second); err != nil {
		fmt.Println("SetConnCheck err", err.Error())
		return
	}

    //Topics need check
    allowList := map[string]interface{}{
        "RECEIVER_A": "aaa",
        "RECEIVER_B": 123.0,
        "RECEIVER_C": struct{}{},
    }

    /*
        Add all topics you need at the beginning when you create nsq producer
        In this way, every time before you publish a topic will be check whether in allow the list
    */
    p, err := config.NewProducer("127.0.0.1:4188", allowList)
    if err != nil {
        //do something
    }

    if err := p.Ping(); err != nil {
        //do something
    }

    //publish success
    msg := "test msg"
    if err := p.Publish("RECEIVER_A", []byte(msg)); err != nil {
        //do something
    }

    //publish fail
    if err := p.Publish("RECEIVER_F", []byte(msg)); err != nil {
        //do something
    }

    hostName, getHostNameErr := os.Hostname()
	if getHostNameErr != nil {
		//do something
	}

    //consumer that will create topic likes `RECEIVER_A-84b7b5897-fm2bg#ephemeral`
    c, err := config.NewConsumerByHost("RECEIVER_A", hostName, "test_receive", true)
	if err != nil {
		//do something
    }

    //do something when your consumer receive message
    c.AddHandler(mynsq.HandlerFunc(func(message *mynsq.Message) error {
		fmt.Println(string(message.Body))
		return nil
    }))

    if err := c.ConnectToNSQLookupd("127.0.0.1:4150"); err != nil {
		fmt.Println("ConnectToNSQLookupd err:", err.Error())
	}
}
```

#### Simple Usage

```go
import (
	"pika.rdtech.vip/genesis-lib/nsq"
)

func init() {
    config := nsq.NewConfig()

    //if you don't want to check topic when each publish, just use nil at NewProducer
    p, err := config.NewProducer("127.0.0.1:4188", nil)
    if err != nil {
        //do something
    }

    msg := "test msg"
    if err := p.Publish("TOPIC", []byte(msg)); err != nil {
        //do something
    }

    done := nsq.NewProducerTransaction()
    if p.PublishAsync("TOPIC", []byte{}, done); err != nil {
        //do something
    }

    //consumer
    c, err := config.NewConsumer("TOPIC", "test_receive")
	if err != nil {
		//do something
    }
    c.AddHandler(mynsq.HandlerFunc(func(message *mynsq.Message) error {
		fmt.Println(string(message.Body))
		return nil
    }))
    if err := c.ConnectToNSQLookupd("127.0.0.1:4150"); err != nil {
		fmt.Println("ConnectToNSQLookupd err:", err.Error())
}
```

## Method

#### Config
- package
- NewConfig
- NewProducer
- NewConsumer
- NewConsumerByHost
- SetDialTimeout
- SetReadTimeout
- SetWriteTimeout
- SetLocalAddr
- SetLookupdPollInterval
- SetLookupdPollJitter
- SetMaxRequeueDelay
- SetDefaultRequeueDelay
- SetMaxBackoffDuration
- SetBackoffMultiplier
- SetMaxAttempts
- SetLowRdyIdleTimeout
- SetLowRdyTimeout
- SetRDYRedistributeInterval
- SetClientID
- SetHostname
- SetUserAgent
- SetHeartbeatInterval
- SetSampleRate
- SetTlsV1
- SetTlsConfig
- SetDeflate
- SetDeflateLevel
- SetSnappy
- SetOutputBufferSize
- SetOutputBufferTimeout
- SetMaxInFlight
- SetMsgTimeout
- SetAuthSecret

#### Producer
- Ping
- Address
- Stop
- Publish
- PublishJSON
- MultiPublish
- MultiPublishJSON
- PublishAsync
- PublishAsyncJSON
- MultiPublishAsync
- MultiPublishAsyncJSON
- DeferredPublish
- DeferredPublishJSON
- DeferredPublishAsync
- DeferredPublishAsyncJSON

#### ProducerTransaction
- NewProducerTransaction