package fingers

import (
	"crypto/sha1"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	s *RPCServer

	// predecessor *Node
	// successor   *Node
	// mu          sync.RWMutex

	// fingerTable fingerTable
	// storage     store
	// transport   Transport

	// lastStablized time.Time

	ID string
}

func NewEchoNode(sock string) (*Node, error) {
	c := Config{
		Socket: sock,
	}

	var o Operations

	return NewNode(&c, o)
}

func NewNode(cfg *Config, logic interface{}) (*Node, error) {
	errVa := cfg.Validate()
	if errVa != nil {
		return nil, errVa
	}

	hostname, errrHo := os.Hostname()
	if errrHo != nil {
		return nil, errrHo
	}

	if len(strings.Trim(hostname, " ")) == 0 {
		return nil, fmt.Errorf("hostname for node listening on socket %s is missing", cfg.Socket)
	}

	s, errNew := NewRPCServer(cfg.Socket, logic)
	if errNew != nil {
		return nil, errNew
	}

	return &Node{
		ID: string(hashWith(sha1.New(), hostname)),
		s:  s,
	}, nil
}

func (n *Node) TalkTo(ipcPort string) error {
	return nil
}

func (n *Node) Start() error {
	return n.s.Start()
}

func (n *Node) Stop() {
	n.s.Stop()
}
