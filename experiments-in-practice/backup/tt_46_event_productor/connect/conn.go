package connect

var CenterConnector *Connector

type Driver interface {
	Build() error
	Detect() error
}

type Connector struct {
	connInfo string
	driver   *Driver
}

func Init(conn string, d *Driver) {
	CenterConnector = NewConnector(conn, d)
}

func NewConnector(conn string, d *Driver) *Connector {
	return &Connector{
		connInfo: conn,
		driver:   d,
	}
}
