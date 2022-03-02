package fingers

import (
	"sync"
	"time"
)

// Transport enables a node to talk to the other nodes in
// the ring.

// https://ipfs.io/ipfs/QmfYeDhGH9bZzihBUDEQbCbTc5k5FZKURMUoUvfmc27BwL/rpc/go_rpc.html

type Transport struct {
	Connections map[string]*Connection
	mu          sync.RWMutex

	maxIdle time.Duration
}

func (t *Transport) cleanIdle() {
	for host, conn := range t.Connections {
		if time.Since(conn.LastActive) > t.maxIdle {
			conn.Close()

			t.mu.Lock()
			delete(t.Connections, host)
			t.mu.Unlock()
		}
	}
}
