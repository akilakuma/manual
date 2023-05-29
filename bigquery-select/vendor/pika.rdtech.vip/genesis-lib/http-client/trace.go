package httpclient

import (
	"crypto/tls"
	"encoding/json"
	"net/http"
	"net/http/httptrace"
	"sync"
	"time"

	"pika.rdtech.vip/genesis-lib/log"
)

//Trace httptrace
type Trace struct {
	URL              string        `json:"url"`
	Method           string        `json:"method"`
	Body             []byte        `json:"body"`
	Header           http.Header   `json:"header"`
	DNSLookup        time.Duration `json:"tc_dns_lookup"`
	TCPConnection    time.Duration `json:"tc_tcp_connection"`
	TLSHandshake     time.Duration `json:"tc_tls_handshake"`
	ServerProcessing time.Duration `json:"tc_server_processing"`
	ContentTransfer  time.Duration `json:"tc_content_transfer"`
	NameLookup       time.Duration `json:"tc_name_lookup"`
	Connect          time.Duration `json:"tc_connect"`
	PreTransfer      time.Duration `json:"tc_pre_transfer"`
	StartTransfer    time.Duration `json:"tc_start_transfer"`
	Total            time.Duration `json:"tc_total"`
	isTLS            bool
	isReused         bool
	lock             *sync.RWMutex
	dnsStart         time.Time
	dnsDone          time.Time
	tcpStart         time.Time
	tcpDone          time.Time
	tlsStart         time.Time
	tlsDone          time.Time
	serverStart      time.Time
	serverDone       time.Time
	transferStart    time.Time
	transferDone     time.Time
}

//Print trace log
func (trace *Trace) Print(message string) {
	log.WithFields(log.Fields{LogPrefix: map[string]interface{}{
		LogTrace: map[string]interface{}{
			"url":                  trace.URL,
			"method":               trace.Method,
			"body":                 jsonEscape(string(trace.Body)),
			"header":               trace.Header,
			"tc_dns_lookup":        trace.DNSLookup.Seconds(),
			"tc_name_lookup":       trace.NameLookup.Seconds(),
			"tc_tcp_connection":    trace.TCPConnection.Seconds(),
			"tc_connect":           trace.Connect.Seconds(),
			"tc_tls_handshake":     trace.TLSHandshake.Seconds(),
			"tc_server_processing": trace.ServerProcessing.Seconds(),
			"tc_pre_transfer":      trace.PreTransfer.Seconds(),
			"tc_start_transfer":    trace.StartTransfer.Seconds(),
			"tc_content_transfer":  trace.ContentTransfer.Seconds(),
			"tc_total":             trace.Total.Seconds(),
		},
	}}).Info(message)
}

func (trace *Trace) createTrace(req *http.Request) *httptrace.ClientTrace {
	trace.lock = &sync.RWMutex{}
	trace.URL = req.Host + req.URL.RequestURI()
	trace.Method = req.Method
	trace.Header = req.Header

	//trace callback
	return &httptrace.ClientTrace{
		DNSStart: func(_ httptrace.DNSStartInfo) {
			trace.lock.Lock()
			trace.dnsStart = time.Now()
			trace.lock.Unlock()
		},
		DNSDone: func(_ httptrace.DNSDoneInfo) {
			trace.lock.Lock()
			trace.dnsDone = time.Now()
			trace.DNSLookup = trace.dnsDone.Sub(trace.dnsStart)
			trace.NameLookup = trace.dnsDone.Sub(trace.dnsStart)
			trace.lock.Unlock()
		},
		ConnectStart: func(_, _ string) {
			trace.lock.Lock()
			trace.tcpStart = time.Now()
			// When connecting to IP (When no DNS lookup)
			if trace.dnsDone.IsZero() {
				trace.dnsStart = trace.tcpStart
				trace.dnsDone = trace.tcpStart
			}
			trace.lock.Unlock()
		},
		ConnectDone: func(_, _ string, _ error) {
			trace.lock.Lock()
			trace.tcpDone = time.Now()
			trace.TCPConnection = trace.tcpDone.Sub(trace.tcpStart)
			trace.Connect = time.Since(trace.dnsStart)
			trace.lock.Unlock()
		},
		TLSHandshakeStart: func() {
			trace.lock.Lock()
			trace.isTLS = true
			trace.tlsStart = time.Now()
			trace.lock.Unlock()
		},
		TLSHandshakeDone: func(_ tls.ConnectionState, _ error) {
			trace.lock.Lock()
			trace.tlsDone = time.Now()
			trace.TLSHandshake = trace.tlsDone.Sub(trace.tlsStart)
			trace.lock.Unlock()
		},
		GotConn: func(i httptrace.GotConnInfo) {
			trace.lock.Lock()
			// Handle when keep alive is used and connection is reused.
			// DNSStart(Done) and ConnectStart(Done) is skipped
			if i.Reused {
				trace.isReused = true
			}
			trace.lock.Unlock()
		},
		WroteRequest: func(_ httptrace.WroteRequestInfo) {
			trace.lock.Lock()
			trace.serverStart = time.Now()
			// When client doesn't use DialContext or using old (before go1.7) `net`
			// pakcage, DNS/TCP/TLS hook is not called.
			if trace.dnsStart.IsZero() && trace.tcpStart.IsZero() {
				trace.dnsStart = trace.serverStart
				trace.dnsDone = trace.serverStart
				trace.tcpStart = trace.serverStart
				trace.tcpDone = trace.serverStart
			}
			trace.PreTransfer = trace.Connect + trace.TLSHandshake
			// When connection is re-used, DNS/TCP/TLS hook is not called.
			if trace.isReused {
				now := trace.serverStart
				trace.dnsStart = now
				trace.dnsDone = now
				trace.tcpStart = now
				trace.tcpDone = now
				trace.tlsStart = now
				trace.tlsDone = now
				trace.TLSHandshake = 0
			}
			if trace.isTLS {
				trace.lock.Unlock()
				return
			}
			// trace.TLSHandshake = trace.tcpDone.Sub(trace.tcpDone)
			trace.lock.Unlock()
		},
		GotFirstResponseByte: func() {
			trace.lock.Lock()
			trace.serverDone = time.Now()
			trace.transferStart = trace.serverDone
			trace.ServerProcessing = trace.serverDone.Sub(trace.serverStart)
			trace.lock.Unlock()
		},
	}
}

func (trace *Trace) endTrace() {
	trace.lock.Lock()
	trace.transferDone = time.Now()
	trace.StartTransfer = trace.PreTransfer + trace.ServerProcessing
	trace.ContentTransfer = trace.transferDone.Sub(trace.transferStart)
	trace.Total = trace.DNSLookup + trace.TCPConnection + trace.TLSHandshake + trace.ServerProcessing + trace.ContentTransfer
	trace.lock.Unlock()
}

func jsonEscape(str string) string {
	b, err := json.Marshal(str)
	if err != nil {
		return str
	}
	// Trim the beginning and trailing " character
	return string(b[1 : len(b)-1])
}
