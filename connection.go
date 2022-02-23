package fingers

import (
	"net/rpc"
	"time"
)

type Connection struct {
	Socket     string
	Conn       *rpc.Client
	LastActive time.Time
}

func (c *Connection) Close() error {
	return c.Conn.Close()
}

func NewConnection() *rpc.Client {
	return &rpc.Client{}
}
