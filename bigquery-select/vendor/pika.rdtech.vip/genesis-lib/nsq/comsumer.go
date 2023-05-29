package nsq

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/buger/jsonparser"
	nsqio "github.com/nsqio/go-nsq"
	"pika.rdtech.vip/eden-lib/inlog"
)

type consumer struct {
	c                  *nsqio.Consumer //original consumer
	config             *nsqio.Config   //original setting
	topic              string          //topic
	channel            string          //channel
	nsqadminAddr       string          //nsqadmin address which can repair consumer connection
	nsqlookupdAddr     string          //nsqlookupd address
	nsqdAddr           string          //nsqd address
	handler            Handler         //consumer handler
	active             bool            //is current consumer actived
	handlerConcurrency int             //amount of concurrency
	connCheckIntervel  time.Duration   //intervel of conn check
	checkSwitch        bool            //switch of repair consumer connection
}

// ConsumerStats represents a snapshot of the state of a Consumer's connections and the messages
// it has seen
type ConsumerStats struct {
	MessagesReceived uint64
	MessagesFinished uint64
	MessagesRequeued uint64
	Connections      int
}

//Consumer is a high-level type to consume from NSQ.
type Consumer interface {
	ConnectToNSQLookupd(addr string) error
	ConnectToNSQD(addr string) error
	AddHandler(handler Handler)
	Stop()
	ConnectToNSQLookupds(addresses []string) error
	ConnectToNSQDs(addresses []string) error
	DisconnectFromNSQD(addr string) error
	DisconnectFromNSQLookupd(addr string) error
	// SetConnCheck(nsqadimAddr string, inervel time.Duration) (err error)
	AddConcurrentHandlers(handler Handler, concurrency int)
	OverWriteLogLevel()
}

// HandlerFunc is a convenience type to avoid having to declare a struct
type HandlerFunc func(message *Message) error

// HandleMessage implements the Handler interface
func (h HandlerFunc) HandleMessage(m *Message) error {
	return h(m)
}

// Handler is the message processing interface for Consumer
type Handler interface {
	HandleMessage(message *Message) error
}

// AddHandler sets the Handler for messages received by this Consumer. This can be called
func (r *consumer) AddHandler(handler Handler) {
	if !r.active {
		r.handler = handler
		r.handlerConcurrency = 1
	}

	oh := func(message *nsqio.Message) error {
		return handler.HandleMessage(&Message{
			Body:        message.Body,
			Timestamp:   message.Timestamp,
			Attempts:    message.Attempts,
			NSQDAddress: message.NSQDAddress,
			msg:         message,
		})
	}
	r.c.AddHandler(nsqio.HandlerFunc(oh))
	r.active = true
}

// AddConcurrentHandlers sets the Handler for messages received by this Consumer.
func (r *consumer) AddConcurrentHandlers(handler Handler, concurrency int) {
	if !r.active {
		r.handler = handler
		r.handlerConcurrency = concurrency
	}

	oh := func(message *nsqio.Message) error {
		return handler.HandleMessage(&Message{
			Body:        message.Body,
			Timestamp:   message.Timestamp,
			Attempts:    message.Attempts,
			NSQDAddress: message.NSQDAddress,
			msg:         message,
		})
	}
	r.c.AddConcurrentHandlers(nsqio.HandlerFunc(oh), concurrency)
	r.active = true
}

//ConnectToNSQD takes a nsqd address to connect directly to.
func (r *consumer) ConnectToNSQD(addr string) error {
	r.nsqdAddr = addr
	connErr := r.c.ConnectToNSQD(addr)
	if err := r.checkConn(r.topic, r.channel); err != nil {
		return err
	}

	return connErr
}

//ConnectToNSQLookupd adds an nsqlookupd address to the list for this Consumer instance.
func (r *consumer) ConnectToNSQLookupd(addr string) error {
	r.nsqlookupdAddr = addr
	connErr := r.c.ConnectToNSQLookupd(addr)
	if err := r.checkConn(r.topic, r.channel); err != nil {
		return err
	}

	return connErr
}

// Stats retrieves the current connection and message statistics for a Consumer
func (r *consumer) Stats() *ConsumerStats {
	c := r.c.Stats()
	return &ConsumerStats{
		MessagesReceived: c.MessagesFinished,
		MessagesFinished: c.MessagesFinished,
		MessagesRequeued: c.MessagesRequeued,
		Connections:      c.Connections,
	}
}

// Stop will initiate a graceful stop of the Consumer (permanent)
//
// NOTE: receive on StopChan to block until this process completes
func (r *consumer) Stop() {
	r.c.Stop()
}

// ConnectToNSQLookupds adds multiple nsqlookupd address to the list for this Consumer instance.
//
// If adding the first address it initiates an HTTP request to discover nsqd
// producers for the configured topic.
//
// A goroutine is spawned to handle continual polling.
func (r *consumer) ConnectToNSQLookupds(addresses []string) error {
	return r.c.ConnectToNSQLookupds(addresses)
}

// ConnectToNSQDs takes multiple nsqd addresses to connect directly to.
//
// It is recommended to use ConnectToNSQLookupd so that topics are discovered
// automatically.  This method is useful when you want to connect to local instance.
func (r *consumer) ConnectToNSQDs(addresses []string) error {
	for _, addr := range addresses {
		err := r.c.ConnectToNSQD(addr)
		if err != nil {
			return err
		}
	}
	return nil
}

// DisconnectFromNSQD closes the connection to and removes the specified
// `nsqd` address from the list
func (r *consumer) DisconnectFromNSQD(addr string) error {
	return r.c.DisconnectFromNSQD(addr)
}

// DisconnectFromNSQLookupd removes the specified `nsqlookupd` address
// from the list used for periodic discovery.
func (r *consumer) DisconnectFromNSQLookupd(addr string) error {
	return r.c.DisconnectFromNSQLookupd(addr)
}

func (r *consumer) checkConn(topic, channel string) error {
	//check switch
	if !r.checkSwitch || r.nsqadminAddr == "" {
		return nil
	}

	//limit
	if r.connCheckIntervel < time.Second {
		r.connCheckIntervel = time.Second
	}

	//new http client
	client := &http.Client{
		Timeout:   5 * time.Second,
		Transport: &http.Transport{MaxConnsPerHost: 100},
	}
	req, reqErr := http.NewRequest("GET", "http://"+r.nsqadminAddr+"/api/topics/"+topic+"/"+channel, nil)
	if reqErr != nil {
		return reqErr
	}

	timer := time.NewTicker(r.connCheckIntervel)
	countFail := 0
	go func() {
		time.Sleep(time.Second) //first round wait
		for {
			<-timer.C //block

			resp, err := client.Do(req)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			data, readErr := ioutil.ReadAll(resp.Body)
			if readErr != nil {
				fmt.Println(err.Error())
				continue
			}
			if err := resp.Body.Close(); err != nil {
				continue
			}
			res, _, _, err := jsonparser.Get(data, "client_count")
			if err != nil {
				fmt.Println(err.Error())
				continue
			}

			switch string(res) {
			case "0":
				countFail++
				//recreate consumer
				if countFail > 2 {
					time.Sleep(r.connCheckIntervel * 2) //waitting other pods
					fmt.Printf("[nsq] consumer topic: %s channel: %s start re-connection\n", topic, channel)
					newConsumer, err := nsqio.NewConsumer(topic, channel, r.config)
					if err != nil {
						fmt.Printf("[nsq] recreate consumer fail. topic: %s, channel: %s\n", topic, channel)
						continue
					}

					//stop old consumer
					if r.c != nil {
						r.c.Stop()
						r.c.DisconnectFromNSQLookupd(r.nsqlookupdAddr)
					}

					//replace by new consumer
					r.c = newConsumer
					r.active = false
					r.AddConcurrentHandlers(r.handler, r.handlerConcurrency)
					if err := r.c.ConnectToNSQLookupd(r.nsqlookupdAddr); err == nil {
						countFail = 0
					}
				}
			default:
				countFail = 0
				continue
			}
		}
	}()

	return nil
}


// SetLogger 強制覆寫log等級成info
func (r *consumer) OverWriteLogLevel() {
	log  := inlog.New()
	nsqLog := NewLogger(log)

	nsqLog.OverWriteLogLevel()

	r.c.SetLogger(nsqLog, 1)
}