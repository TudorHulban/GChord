package fingers

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"net"
	"net/rpc"
	"os"
	"strings"
	"sync"
	"time"
)

type Config struct {
	Hash     hasher
	HashSize int

	Socket string

	StabilizeMin time.Duration
	StabilizeMax time.Duration

	Timeout time.Duration
	MaxIdle time.Duration
}

type Node struct {
	cfg *Config

	predecessor *Node
	successor   *Node
	mu          sync.RWMutex

	fingerTable fingerTable
	storage     store
	transport   Transport

	chShutdown    chan struct{} // TODO: assess best type
	lastStablized time.Time

	ID string
}

func (c *Config) Validate() error {
	return nil
}

func NewNode(cfg *Config) (*Node, error) {
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

	return &Node{
		ID:         string(hashWith(sha1.New(), hostname)),
		chShutdown: make(chan struct{}),
	}, nil
}

func (n *Node) Start(o *Operations) error {
	errRe := rpc.Register(o)
	if errRe != nil {
		return errRe
	}

	tcpAddr, errRes := net.ResolveTCPAddr("tcp", n.cfg.Socket)
	if errRes != nil {
		return errRes
	}

	if tcpAddr == nil {
		return errors.New("tcpAddr is nil")
	}

	// _, errLis := net.ListenTCP("tcp", tcpAddr)
	// if errLis != nil {
	// 	return errLis
	// }

	// go func() {
	// 	fmt.Printf("listening on port %s\n", n.cfg.Socket)

	// 	select {
	// 	case sign := <-n.chShutdown:
	// 		{
	// 			fmt.Printf("\nshutting node on received: %s.\n", sign)
	// 			close(n.chShutdown)
	// 		}

	// 	default:
	// 		{
	// 			// rpc.Accept(listener)
	// 		}
	// 	}
	// }()

	return nil
}

func (n *Node) Stop() {
	n.chShutdown <- struct{}{}
}
