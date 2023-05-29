package nsq

import (
	"io"
	"time"

	nsqio "github.com/nsqio/go-nsq"
)

//MsgIDLength The number of bytes for a Message.ID
const MsgIDLength = nsqio.MsgIDLength

// MessageID is the ASCII encoded hexadecimal message ID
type MessageID [MsgIDLength]byte

// MessageDelegate is an interface of methods that are used as
// callbacks in Message
type MessageDelegate interface {
	// OnFinish is called when the Finish() method
	// is triggered on the Message
	OnFinish(*Message)

	// OnRequeue is called when the Requeue() method
	// is triggered on the Message
	OnRequeue(m *Message, delay time.Duration, backoff bool)

	// OnTouch is called when the Touch() method
	// is triggered on the Message
	OnTouch(*Message)
}

// Message is the fundamental data type containing
// the id, body, and metadata
type Message struct {
	ID                   MessageID
	Body                 []byte
	Timestamp            int64
	Attempts             uint16
	NSQDAddress          string
	Delegate             MessageDelegate
	autoResponseDisabled int32
	responded            int32
	msg                  *nsqio.Message
}

// DisableAutoResponse disables the automatic response
func (m *Message) DisableAutoResponse() {
	m.msg.DisableAutoResponse()
}

// IsAutoResponseDisabled indicates whether or not this message
// will be responded to automatically
func (m *Message) IsAutoResponseDisabled() bool {
	return m.msg.IsAutoResponseDisabled()
}

// HasResponded indicates whether or not this message has been responded to
func (m *Message) HasResponded() bool {
	return m.msg.HasResponded()
}

// Finish sends a FIN command to the nsqd which
// sent this message
func (m *Message) Finish() {
	m.msg.Finish()
}

// Touch sends a TOUCH command to the nsqd which
// sent this message
func (m *Message) Touch() {
	m.msg.Touch()
}

// Requeue sends a REQ command to the nsqd which
// sent this message, using the supplied delay.
//
// A delay of -1 will automatically calculate
// based on the number of attempts and the
// configured default_requeue_delay
func (m *Message) Requeue(delay time.Duration) {
	m.msg.Requeue(delay)
}

// RequeueWithoutBackoff sends a REQ command to the nsqd which
// sent this message, using the supplied delay.
//
// Notably, using this method to respond does not trigger a backoff
// event on the configured Delegate.
func (m *Message) RequeueWithoutBackoff(delay time.Duration) {
	m.RequeueWithoutBackoff(delay)
}

// WriteTo implements the WriterTo interface and serializes
// the message into the supplied producer.
//
// It is suggested that the target Writer is buffered to
// avoid performing many system calls.
func (m *Message) WriteTo(w io.Writer) (int64, error) {
	return m.msg.WriteTo(w)
}
