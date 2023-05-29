package nsq

import (
	"crypto/tls"
	"errors"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"

	"pika.rdtech.vip/eden-lib/inlog"
	nsqio "github.com/nsqio/go-nsq"
)

// Config is a struct of NSQ options
type Config struct {
	setting           *nsqio.Config  //original setting
	logger            *inlog.Logger  //logger
	logLevel          nsqio.LogLevel //log level
	checkSwitch       bool           //switch of repair consumer connection
	nsqadminAddr      string         //nsqadmin address which can repair consumer connection
	connCheckIntervel time.Duration  //intervel of conn check
}

//NewConfig returns a new default nsq configuration.
func NewConfig() *Config {
	return &Config{setting: nsqio.NewConfig()}
}

//NewConsumer creates a new instance of Consumer for the specified topic/channel
func (c *Config) NewConsumer(topic string, channel string) (Consumer, error) {
	original, err := nsqio.NewConsumer(topic, channel, c.setting)
	var nsqLogger Logger

	if c.logger != nil {
		nsqLogger = NewLogger(c.logger)
	} else {
		l := inlog.New()
		nsqLogger = NewLogger(l)
	}

	original.SetLogger(nsqLogger, c.logLevel)
	return &consumer{c: original, config: c.setting, topic: topic, channel: channel, checkSwitch: c.checkSwitch, nsqadminAddr: c.nsqadminAddr, connCheckIntervel: c.connCheckIntervel}, err
}

//NewConsumerByHost creates a new instance of Consumer for the specified host/channel
func (c *Config) NewConsumerByHost(topic, host, channel string, ephemeral bool) (Consumer, error) {
	if host == "" {
		return nil, errors.New("empty host")
	}

	hostChannel := channel
	hostChannel = host + channel
	if ephemeral {
		hostChannel += "#ephemeral"
	}

	original, err := nsqio.NewConsumer(topic, hostChannel, c.setting)
	var nsqLogger Logger

	if c.logger != nil {
		nsqLogger = NewLogger(c.logger)
	} else {
		l := inlog.New()
		nsqLogger = NewLogger(l)
	}

	original.SetLogger(nsqLogger, c.logLevel)
	return &consumer{c: original, config: c.setting, topic: topic, channel: channel, checkSwitch: c.checkSwitch, nsqadminAddr: c.nsqadminAddr, connCheckIntervel: c.connCheckIntervel}, err
}

//NewProducer returns an instance of Producer for the specified address
func (c *Config) NewProducer(addr string, topics map[string]interface{}) (Producer, error) {
	var (
		check   bool
		syncMap sync.Map
	)

	if addr == "" {
		return nil, errors.New("nsqd addr empty")
	}

	if topics != nil && len(topics) > 0 {
		check = true
		for k, v := range topics {
			syncMap.Store(k, v)
			if err := createTopic(addr, k); err != nil {
				return nil, err
			}
		}
	}

	p, err := nsqio.NewProducer(addr, c.setting)
	var nsqLogger Logger

	if c.logger != nil {
		nsqLogger = NewLogger(c.logger)
	} else {
		l := inlog.New()
		nsqLogger = NewLogger(l)
	}

	p.SetLogger(nsqLogger, c.logLevel)
	return &producer{p: p, topics: &syncMap, topicCheck: check}, err
}

func createTopic(nsqdAddr, topic string) error {
	client := &http.Client{}
	addr := strings.Split(nsqdAddr, ":")

	req, err := http.NewRequest("POST", "http://"+addr[0]+":4151/topic/create?topic="+topic, nil)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			return
		}
	}()

	if _, err := ioutil.ReadAll(resp.Body); err != nil {
		return err
	}

	return nil
}

// SetDialTimeout SetDialTimeout
func (c *Config) SetDialTimeout(t time.Duration) {
	c.setting.DialTimeout = t
}

// SetReadTimeout SetReadTimeout
func (c *Config) SetReadTimeout(t time.Duration) {
	c.setting.ReadTimeout = t
}

// SetWriteTimeout SetWriteTimeout
func (c *Config) SetWriteTimeout(t time.Duration) {
	c.setting.WriteTimeout = t
}

// SetLocalAddr SetLocalAddr
func (c *Config) SetLocalAddr(addr net.Addr) {
	c.setting.LocalAddr = addr
}

// SetLookupdPollInterval SetLookupdPollInterval
func (c *Config) SetLookupdPollInterval(t time.Duration) {
	c.setting.LookupdPollInterval = t
}

// SetLookupdPollJitter SetLookupdPollJitter
func (c *Config) SetLookupdPollJitter(jitter float64) {
	c.setting.LookupdPollJitter = jitter
}

// SetMaxRequeueDelay SetMaxRequeueDelay
func (c *Config) SetMaxRequeueDelay(t time.Duration) {
	c.setting.MaxRequeueDelay = t
}

// SetDefaultRequeueDelay SetDefaultRequeueDelay
func (c *Config) SetDefaultRequeueDelay(t time.Duration) {
	c.setting.DefaultRequeueDelay = t
}

// SetMaxBackoffDuration SetMaxBackoffDuration
func (c *Config) SetMaxBackoffDuration(t time.Duration) {
	c.setting.MaxBackoffDuration = t
}

// SetBackoffMultiplier SetBackoffMultiplier
func (c *Config) SetBackoffMultiplier(t time.Duration) {
	c.setting.BackoffMultiplier = t
}

// SetMaxAttempts SetMaxAttempts
func (c *Config) SetMaxAttempts(times uint16) {
	c.setting.MaxAttempts = times
}

// SetLowRdyIdleTimeout SetLowRdyIdleTimeout
func (c *Config) SetLowRdyIdleTimeout(t time.Duration) {
	c.setting.LowRdyIdleTimeout = t
}

// SetLowRdyTimeout SetLowRdyTimeout
func (c *Config) SetLowRdyTimeout(t time.Duration) {
	c.setting.LowRdyTimeout = t
}

// SetRDYRedistributeInterval SetRDYRedistributeInterval
func (c *Config) SetRDYRedistributeInterval(t time.Duration) {
	c.setting.RDYRedistributeInterval = t
}

// SetClientID SetClientID
func (c *Config) SetClientID(id string) {
	c.setting.ClientID = id
}

// SetHostname SetHostname
func (c *Config) SetHostname(name string) {
	c.setting.Hostname = name
}

// SetUserAgent SetUserAgent
func (c *Config) SetUserAgent(agent string) {
	c.setting.UserAgent = agent
}

// SetHeartbeatInterval SetHeartbeatInterval
func (c *Config) SetHeartbeatInterval(t time.Duration) {
	c.setting.HeartbeatInterval = t
}

// SetSampleRate SetSampleRate
func (c *Config) SetSampleRate(rate int32) {
	c.setting.SampleRate = rate
}

// SetTLSV1 SetTlsV1
func (c *Config) SetTLSV1(open bool) {
	c.setting.TlsV1 = open
}

// SetTLSConfig SetTlsConfig
func (c *Config) SetTLSConfig(tlsC *tls.Config) {
	c.setting.TlsConfig = tlsC
}

// SetDeflate SetDeflate
func (c *Config) SetDeflate(deflate bool) {
	c.setting.Deflate = deflate
}

// SetDeflateLevel SetDeflateLevel
func (c *Config) SetDeflateLevel(level int) {
	c.setting.DeflateLevel = level
}

// SetSnappy SetSnappy
func (c *Config) SetSnappy(snappy bool) {
	c.setting.Snappy = snappy
}

// SetOutputBufferSize SetOutputBufferSize
func (c *Config) SetOutputBufferSize(b int64) {
	c.setting.OutputBufferSize = b
}

// SetOutputBufferTimeout SetOutputBufferTimeout
func (c *Config) SetOutputBufferTimeout(t time.Duration) {
	c.setting.OutputBufferTimeout = t
}

// SetMaxInFlight SetMaxInFlight
func (c *Config) SetMaxInFlight(max int) {
	c.setting.MaxInFlight = max
}

// SetMsgTimeout SetMsgTimeout
func (c *Config) SetMsgTimeout(t time.Duration) {
	c.setting.MsgTimeout = t
}

// SetAuthSecret SetAuthSecret
func (c *Config) SetAuthSecret(auth string) {
	c.setting.AuthSecret = auth
}

//SetLogger SetLogger
func (c *Config) SetLogger(l *inlog.Logger) {
	c.logger = l
}

//SetLogLevel SetLogLevel
func (c *Config) SetLogLevel(lv InLogLevel) {
	var nsqlv nsqio.LogLevel
	switch lv {
	case LogLevelDebug:
		nsqlv = nsqio.LogLevelDebug
	case LogLevelInfo:
		nsqlv = nsqio.LogLevelInfo
	case LogLevelWarning:
		nsqlv = nsqio.LogLevelWarning
	case LogLevelError:
		nsqlv = nsqio.LogLevelError
	default:
		nsqlv = nsqio.LogLevelDebug
	}

	c.logLevel = nsqlv
}

//SetConnCheck Open sonsumer connection check and repair
func (c *Config) SetConnCheck(nsqadimAddr string, inervel time.Duration) (err error) {
	if nsqadimAddr == "" {
		return errors.New("empty nsqadmin address")
	}

	if inervel < time.Second {
		return errors.New("intervel should not smaller than 1s")
	}

	c.checkSwitch = true
	c.nsqadminAddr = nsqadimAddr
	c.connCheckIntervel = inervel
	return
}
