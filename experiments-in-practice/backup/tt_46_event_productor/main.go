package main

import (
	"Golang-Learning/tt_46_event_productor/conn"
	"log"
	"sync"
	"time"
)

// EventController 事件控制
type EventController struct {
	id    int
	event Event
}

type MyQQDriver struct{}

func (mqq *MyQQDriver) Build() {

}

func (mqq *MyQQDriver) Detect() {

}

func main() {

	conn.Init("192.168.2.1", MyQQDriver)
	conn.CenterConnector.Build()

	em := NewEM()

	event1 := &Event{
		Name:     "e1",
		Period:   5 * time.Second,
		IsRepeat: true,
		Action:   action1,
	}

	event2 := &Event{
		Name:     "e2",
		Period:   7 * time.Second,
		IsRepeat: true,
		Action:   action2,
	}

	event3 := &Event{
		Name:     "e3",
		Period:   10 * time.Second,
		IsRepeat: true,
		Action:   action3,
	}

	event4 := &Event{
		Name:     "e4",
		Period:   2 * time.Second,
		IsRepeat: false,
		Action:   action4,
	}

	em.PushEvent(event1)
	em.PushEvent(event2)
	em.PushEvent(event3)
	em.PushEvent(event4)

	go em.Running()

	for {
		time.Sleep(30 * time.Second)
		go em.Close()
	}
}

func action1() error {
	log.Println("我是一隻貓，喵喵！")
	time.Sleep(1 * time.Second)
	return nil
}

func action2() error {
	log.Println("浣熊神敎，只能覺得浣熊最可愛")
	time.Sleep(2 * time.Second)
	return nil
}

func action3() error {
	log.Println("今天天氣真好")
	time.Sleep(3 * time.Second)
	return nil
}

func action4() error {
	log.Println("大家好")
	time.Sleep(1 * time.Second)
	return nil
}

// EventManager 事件管理
type EventManager struct {
	jobQueue chan *Event
	exitMsg  chan struct{}
}

// Event 事件資料
type Event struct {
	Name     string
	Period   time.Duration
	IsRepeat bool
	Action   func() error
}

// NewEM 新的管理器實體
func NewEM() *EventManager {
	return &EventManager{
		jobQueue: make(chan *Event, 100),
		exitMsg:  make(chan struct{}),
	}
}

// PushEvent 加入event
func (em *EventManager) PushEvent(event *Event) {
	em.jobQueue <- event
}

// PopEvent 拿出event
func (em *EventManager) PopEvent() *Event {
	event, ok := <-em.jobQueue
	if ok {
		return event
	}

	return nil
}

// Running 啟動執行
func (em *EventManager) Running() {

	log.Println("start running")

	// 首先拿出事件
	// 查看事件的設定時間
	// 根據這個時間設定一個timer
	event := em.PopEvent()
	t := event.Period
	eventTimer := time.NewTimer(t)

	for {
		select {
		// timer 到了
		case <-eventTimer.C:

			log.Println("time arrive")

			eventTimer.Stop()

			// 執行event
			go event.Action()

			// 如果需要重複的話，塞回channel末端
			if event.IsRepeat {
				em.PushEvent(event)
			}

			event = em.PopEvent()
			t = event.Period
			eventTimer = time.NewTimer(t)
		case <-em.exitMsg:

			log.Println("leave")
			return
		}
	}
}

// Close 啟動執行
func (em *EventManager) Close() {
	em.exitMsg <- struct{}{}
}

// NewEvent 新的Event
func NewEvent(per time.Duration, rp bool, act func() error) *Event {
	return &Event{
		Period:   per,
		IsRepeat: rp,
		Action:   act,
	}
}

// SetName 設定事件名稱
func (e *Event) SetName(n string) {
	e.Name = n
}

type Message struct {
	RWLocker *sync.RWMutex
	MMap     map[int]string
}

type MMController struct {
	m *Message
}

func NewMessage() *Message {
	return &Message{
		RWLocker: new(sync.RWMutex),
		MMap:     make(map[int]string, 0),
	}
}

func NewMMControl() *MMController {
	return &MMController{
		m: NewMessage(),
	}
}
