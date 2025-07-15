package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"net"
	"net/http"
	"onion/global"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	var r *gin.Engine

	r = gin.Default()
	meow := r.Group("/")

	meow.GET("hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, "hello man !")
		return
	})
	r.Run(":2266")

	// 建立一個伺服器
	// server := CreateServer(r, "localhost:2266")

	// 啟動伺服器監聽
	// SignalListenAndServe(server, "localhost:2266")
}

func test1() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`hello world`))
	})
	http.ListenAndServe(":3000", nil)
}

// CreateServer 建立伺服器
func CreateServer(r *gin.Engine, port string) *http.Server {
	// 建立 Server
	server := &http.Server{
		Addr:         port,
		Handler:      r,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		// MaxHeaderBytes: 1 << 20,
	}

	return server
}

// SignalListenAndServe 開啟Server & 系統信號監聽
func SignalListenAndServe(server *http.Server, port string) {
	defer func() {
		if err := recover(); err != nil {
			errMessage := fmt.Sprintf("❌  Server 發生意外 Error: %v ❌", err)
			global.WriteLog("ERROR", errMessage)
			global.NotifyEngineer(errMessage)
		}
	}()

	global.SetupGracefulSignal()

	l, err := net.Listen("tcp", port)
	if err != nil {
		errMessage := fmt.Sprintf("❌  Server 建立監聽 Error: %v ❌", err)
		global.WriteLog("ERROR", errMessage)
		global.NotifyEngineer(errMessage)
		return
	}

	dl := NewDozListner(l, 0)

	global.WriteLog("INFO", fmt.Sprintf("🐠  Server 開始服務! %s 🐠", l.Addr().String()))
	defer global.WriteLog("INFO", "🔥  Server 結束服務!🔥")

	go server.Serve(dl)
	// go server.Serve(l)

	errCh := make(chan error, 2)
	go func() {
		receivedSignal := <-global.GracefulDown()
		global.WriteLog("INFO", fmt.Sprintf("🎃  Server 接受訊號 <- %v 🎃", receivedSignal))
		dl.Close()
		dl.Done()
		// l.Close()
		errCh <- nil
	}()

	// 等待結束
	select {
	case err := <-errCh:
		if err != nil {
			return
		}
	}
}

// DozListener 監聽
type DozListener struct {
	sync.Once
	net.Listener
	buf   chan struct{}
	out   chan struct{}
	errCh chan error
	sig   chan os.Signal
}

// DozConn 連線
type DozConn struct {
	net.Conn
	onClosed func()
}

// Close 關閉連線
func (conn *DozConn) Close() error {
	// log.Println("<=== 連線關閉...", conn.Conn.RemoteAddr().String())
	err := conn.Conn.Close()
	conn.onClosed()
	return err
}

func (l *DozListener) buffOut() {
	select {
	case <-l.errCh:
		l.out <- <-l.buf
	default:
		<-l.buf
	}
}

func (l *DozListener) onlyDoWithBuffer(do func() error) error {
	if l.buf != nil {
		return do()
	}

	return nil
}

// Accept 接收連線
func (l *DozListener) Accept() (net.Conn, error) {

	ttS := time.Now()

	err := l.onlyDoWithBuffer(func() error {
		// 如果現在Buffer滿了或關閉了，不接收連線
		select {
		case <-l.errCh:
			return errors.New("DozListener Closed")
		case l.buf <- struct{}{}:
			// log.Println("等待連線...")
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	ttE := time.Now().Sub(ttS)
	if ttE > time.Duration(10*time.Millisecond) {
		log.Println("🗝 🗝 🗝", ttE)
	}

	conn, err := l.Listener.Accept()
	ttS2 := time.Now()
	if err != nil {
		doErr := l.onlyDoWithBuffer(func() error {
			l.buffOut()
			select {
			case <-l.errCh:
				return errors.New("DozListener Closed")
			default:
			}
			return nil
		})

		if doErr != nil {
			err = doErr
		}

		log.Println("接收連線Error...", err)
		return nil, err
	}
	// log.Println("===> 接收連線...", conn.RemoteAddr().String())
	// log.Printf("⭐  目前Goroutine數量: %d, 連線數量: %d", runtime.NumGoroutine(), len(l.buf))

	ttE2 := time.Now().Sub(ttS2)
	if ttE2 > time.Duration(10*time.Millisecond) {
		log.Println("⚰ ⚰ ⚰", ttE)
	}

	return &DozConn{
		Conn:     conn,
		onClosed: l.buffOut,
	}, nil
}

// Close 關閉監聽
func (l *DozListener) Close() error {
	var err error
	l.Do(func() {
		l.onlyDoWithBuffer(func() error {
			l.out = make(chan struct{}, cap(l.buf))
			// close(l.buf)
			close(l.errCh)
			return nil
		})
		err = l.Listener.Close()
		log.Println("關閉監聽...", err)
	})
	return err
}

// Done 等待連線關閉
func (l *DozListener) Done() error {
	err := l.onlyDoWithBuffer(func() error {
		var notOver bool
		var count int

		count = len(l.buf)
		notOver = count > 0

		for notOver {
			log.Printf("🎯  還有%d條連線，等待關閉...", count)
			select {
			case <-l.out:
				count = len(l.buf)
				if count == 0 {
					log.Println("連線已經清空")
					notOver = false
					break
				}
			case <-l.sig:
				os.Exit(127)
			}
		}

		return nil
	})

	return err
}

// NewDozListner 建立新的監聽
func NewDozListner(l net.Listener, poolSize int) *DozListener {
	dl := &DozListener{
		Listener: l,
		errCh:    make(chan error),
	}

	sig := make(chan os.Signal, 0)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	dl.sig = sig

	if poolSize > 0 {
		dl.buf = make(chan struct{}, poolSize)
	} else {
		dl.buf = make(chan struct{}, math.MaxInt64)
	}
	return dl
}
