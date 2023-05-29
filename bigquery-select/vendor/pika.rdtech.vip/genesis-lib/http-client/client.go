package httpclient

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/syhlion/greq"
	goworker "github.com/syhlion/requestwork.v2"
	"pika.rdtech.vip/genesis-lib/log"
)

//const const
const (
	MethodGET            = "GET"
	MethodPOST           = "POST"
	MethodPUT            = "PUT"
	MethodDELETE         = "DELETE"
	DefaultMaxConn   int = 100
	DefaultQueueConn int = 100
	DefaultTimeOut       = 60 * time.Second
	DefaultAddr          = "127.0.0.1"
	DefaultSSL           = "http"
	LogPrefix            = "http-client"
	LogError             = "error"
	LogTrace             = "trace"
)

//Methods Basic usage that you have to imprement
type Methods interface {
	Get(ctx context.Context, uri string, param map[string]string) (resp Response)
	Post(ctx context.Context, uri string, body []byte) (resp Response)
	Put(ctx context.Context, uri string, body []byte) (resp Response)
	Delete(ctx context.Context, uri string, body []byte) (resp Response)
	GetWithHeader(ctx context.Context, uri string, param, header, cookie map[string]string) (resp Response)
	PostWithHeader(ctx context.Context, uri string, body []byte, header, cookie map[string]string) (resp Response)
	PutWithHeader(ctx context.Context, uri string, body []byte, header, cookie map[string]string) (resp Response)
	DeleteWithHeader(ctx context.Context, uri string, body []byte, header, cookie map[string]string) (resp Response)
	GetTimeOut() (timeout time.Duration)
	GetHeader(key string) (value string)
	GetHost() string
	GetSSL() string
	GetAddr() (addr string)
	GetWithBody(ctx context.Context, uri string, param map[string]string, body []byte, header, cookie map[string]string) (resp Response)
}

//Conn Client connection info
type Conn struct {
	debug       bool              //open debug mode
	newClientAt int64             //client instanst created time
	maxConn     int               //max connection
	ssl         string            //http|https
	host        string            //www.wtf.tw
	ipAddr      string            //127.0.0.1
	timeout     time.Duration     //request connection timeout
	header      map[string]string //header [key]value
	lock        *sync.RWMutex     //goroutine lock
	worker      *goworker.Worker  //conn exec worker
	ctx         context.Context   //context
	greqConn    *greq.Client
}

//Body Body
type Body []byte

//Response Response
type Response struct {
	Header  http.Header
	Trace   *Trace
	Cookies []*http.Cookie
	Body    Body
	Code    int
	Err     error
}

func init() {
	log.SetFormatter(&log.JSONFormatter{})
}

//New Create a new client connection setting
func New(ctx context.Context, threads int, opts ...Option) Methods {
	conn := Conn{
		newClientAt: time.Now().Unix(),
		header:      make(map[string]string),
		maxConn:     threads,
		timeout:     DefaultTimeOut,
		ipAddr:      DefaultAddr,
		ssl:         DefaultSSL,

		lock: &sync.RWMutex{},
		ctx:  ctx,
	}

	conn.worker = goworker.New(threads)

	for _, o := range opts {
		o.apply(&conn)
	}

	conn.greqConn = greq.New(conn.worker, conn.timeout, false)

	if conn.host != "" {
		conn.greqConn.SetHost(conn.host)
	}

	for k, v := range conn.header {
		conn.greqConn.SetHeader(k, v)
	}

	return &conn
}

//Get Method GetWithCtx
func (c *Conn) Get(ctx context.Context, uri string, param map[string]string) (resp Response) {
	if ok, r := uriCheck(uri); !ok {
		return r
	}

	//req
	url := joinURI(c.formateURL(uri), param)
	req, err := http.NewRequest(MethodGET, url, nil)
	if err != nil {
		resp.Err = err
		return
	}

	//header
	c.headerHandeler(req)

	//resp
	return c.requestHandler(req, nil)
}

//GetWithHeader Method GetWithHeaderCtx
func (c *Conn) GetWithHeader(ctx context.Context, uri string, param, header, cookie map[string]string) (resp Response) {
	if ok, r := uriCheck(uri); !ok {
		return r
	}

	//req
	url := joinURI(c.formateURL(uri), param)
	req, err := http.NewRequest(MethodGET, url, nil)
	if err != nil {
		resp.Err = err
		return
	}

	//header
	c.onceHeaderHandeler(req, header, cookie)

	//resp
	return c.requestHandler(req, nil)
}

//GetWithBody Method Get with payload
func (c *Conn) GetWithBody(ctx context.Context, uri string, param map[string]string, body []byte, header, cookie map[string]string) (resp Response) {
	if ok, r := uriCheck(uri); !ok {
		return r
	}

	//req
	url := joinURI(c.formateURL(uri), param)
	req, err := http.NewRequest(MethodGET, url, bytes.NewReader(body))
	if err != nil {
		resp.Err = err
		return
	}

	//header
	c.onceHeaderHandeler(req, header, cookie)

	//resp
	return c.requestHandler(req, body)
}

//Post Method PostWithCtx
func (c *Conn) Post(ctx context.Context, uri string, body []byte) (resp Response) {
	if ok, r := uriCheck(uri); !ok {
		return r
	}

	//req
	req, err := http.NewRequest(MethodPOST, c.formateURL(uri), bytes.NewReader(body))
	if err != nil {
		resp.Err = err
		return
	}

	//header
	c.headerHandeler(req)

	//resp
	return c.requestHandler(req, body)
}

//PostWithHeader Method PostWithHeaderCtx
func (c *Conn) PostWithHeader(ctx context.Context, uri string, body []byte, header, cookie map[string]string) (resp Response) {
	if ok, r := uriCheck(uri); !ok {
		return r
	}

	//req
	req, err := http.NewRequest(MethodPOST, c.formateURL(uri), bytes.NewReader(body))
	if err != nil {
		resp.Err = err
		return
	}

	//header
	c.onceHeaderHandeler(req, header, cookie)

	//resp
	return c.requestHandler(req, body)
}

//Put Method PutWithCtx
func (c *Conn) Put(ctx context.Context, uri string, body []byte) (resp Response) {
	if ok, r := uriCheck(uri); !ok {
		return r
	}

	//req
	req, err := http.NewRequest(MethodPUT, c.formateURL(uri), bytes.NewReader(body))
	if err != nil {
		resp.Err = err
		return
	}

	//header
	c.headerHandeler(req)

	//resp
	return c.requestHandler(req, body)
}

//PutWithHeader Method PutWithHeaderCtx
func (c *Conn) PutWithHeader(ctx context.Context, uri string, body []byte, header, cookie map[string]string) (resp Response) {
	if ok, r := uriCheck(uri); !ok {
		return r
	}

	//req
	req, err := http.NewRequest(MethodPUT, c.formateURL(uri), bytes.NewReader(body))
	if err != nil {
		resp.Err = err
		return
	}

	//header
	c.onceHeaderHandeler(req, header, cookie)

	//resp
	return c.requestHandler(req, body)
}

//Delete Method DeleteWithCtx
func (c *Conn) Delete(ctx context.Context, uri string, body []byte) (resp Response) {
	if ok, r := uriCheck(uri); !ok {
		return r
	}

	//req
	req, err := http.NewRequest(MethodDELETE, c.formateURL(uri), bytes.NewReader(body))
	if err != nil {
		resp.Err = err
		return
	}

	//header
	c.headerHandeler(req)

	//resp
	return c.requestHandler(req, body)
}

//DeleteWithHeader Method DeleteWithHeaderCtx
func (c *Conn) DeleteWithHeader(ctx context.Context, uri string, body []byte, header, cookie map[string]string) (resp Response) {
	if ok, r := uriCheck(uri); !ok {
		return r
	}

	//req
	req, err := http.NewRequest(MethodDELETE, c.formateURL(uri), bytes.NewReader(body))
	if err != nil {
		resp.Err = err
		return
	}

	//header
	c.onceHeaderHandeler(req, header, cookie)

	//resp
	return c.requestHandler(req, body)
}

//GetHeader GetHeader
func (c *Conn) GetHeader(key string) (value string) {
	c.lock.RLock()
	if v, ok := c.header[key]; ok {
		value = v
	}
	c.lock.RUnlock()
	return
}

//String 處理 response body
func (r Body) String() (strBody string) {
	strBody = string(r)

	return
}

//GetTimeOut GetTimeOut
func (c *Conn) GetTimeOut() (timeout time.Duration) {
	return c.timeout
}

//GetHost GetHost
func (c *Conn) GetHost() string {
	return c.host
}

//GetSSL GetSSL
func (c *Conn) GetSSL() string {
	return c.ssl
}

//GetAddr GetAddr
func (c *Conn) GetAddr() (addr string) {
	return c.ipAddr
}

func uriCheck(uri string) (ok bool, r Response) {
	if uri == "" {
		r.Err = errors.New("URI nil")
		return
	}
	ok = true
	return
}

func (c *Conn) formateURL(uri string) (url string) {
	addr := c.ipAddr
	if c.ipAddr == "127.0.0.1" && c.host != "" {
		addr = c.host
	}

	if !strings.Contains(uri, "http:") && !strings.Contains(uri, "https:") {
		url = c.ssl + "://" + addr + uri
	} else {
		url = uri
	}

	return
}

func joinURI(uri string, param map[string]string) string {
	// Set query param
	if len(param) > 0 {
		q := url.Values{}
		for key, val := range param {
			q.Add(key, val)
		}
		uri = uri + "?" + q.Encode()
	}

	return uri
}

//Add Conn.header to http.NewRequest
func (c *Conn) headerHandeler(req *http.Request) {
	c.lock.RLock()
	for key, value := range c.header {
		req.Header.Set(key, value)
	}
	c.lock.RUnlock()
}

//Add Conn.header to http.NewRequest
func (c *Conn) onceHeaderHandeler(req *http.Request, header, cookies map[string]string) {
	if len(header) > 0 {
		for key, value := range header {
			req.Header.Set(key, value)
		}
	}

	if len(cookies) > 0 {
		for key, value := range cookies {
			req.AddCookie(&http.Cookie{
				Name:  key,
				Value: value,
			})
		}
	}
}

func (c *Conn) requestHandler(req *http.Request, body []byte) (resp Response) {
	//http trace
	trace := &Trace{Body: body}

	//resp
	resp.Body, resp.Code, resp.Err = c.greqConn.ResolveTraceRequest(req, trace.createTrace(req))
	trace.endTrace()
	resp.Trace = trace
	return
}
