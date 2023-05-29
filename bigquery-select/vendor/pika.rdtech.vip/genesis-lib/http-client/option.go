package httpclient

import (
	"net/http"
	"time"
)

//Option Client setting
type Option interface {
	apply(*Conn)
}

type funcOption struct {
	f func(*Conn)
}

func (fo *funcOption) apply(c *Conn) {
	fo.f(c)
}

func newFuncOption(f func(*Conn)) *funcOption {
	return &funcOption{
		f: f,
	}
}

//SetHTTPTransport SetHTTPTransport
func SetHTTPTransport(trs *http.Transport) Option {
	return newFuncOption(func(c *Conn) {
		c.worker.SetTransport(trs)
	})
}

//SetAddr SetAddr
func SetAddr(ipAddr string) Option {
	return newFuncOption(func(c *Conn) {
		c.ipAddr = ipAddr
	})
}

//SetHeader SetHeader
func SetHeader(header map[string]string) Option {
	return newFuncOption(func(c *Conn) {
		c.lock.Lock()
		for k, v := range header {
			c.header[k] = v
		}
		c.lock.Unlock()
	})
}

//SetTimeOut SetTimeOut
func SetTimeOut(timeout time.Duration) Option {
	return newFuncOption(func(c *Conn) {
		c.timeout = timeout
	})
}

//SetHost SetHost
func SetHost(host string) Option {
	return newFuncOption(func(c *Conn) {
		c.host = host
	})
}

// SetSSL sets URI ssl, i.e. http, https, ftp, etc.
func SetSSL(ssl string) Option {
	return newFuncOption(func(c *Conn) {
		c.ssl = ssl
	})
}

//SetDebug open debug trace log of http connection transfer
func SetDebug(open bool) Option {
	return newFuncOption(func(c *Conn) {
		c.debug = open
	})
}
