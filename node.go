package fingers

import (
	"fmt"
	"hash"
	"net"
	"net/rpc"
	"sync"
	"time"
)

type hasher func() hash.Hash
type Socket string

type Config struct {
	Hash     hasher
	HashSize int

	Socket Socket

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
}

func NewNode(cfg *Config) (*Node, error) {
	// TODO: validation of config

	return &Node{
		chShutdown: make(chan struct{}),
	}, nil
}

func (n *Node) Start(o *Operations) error {
	rpc.Register(&o)

	tcpAddr, errRes := net.ResolveTCPAddr("tcp", string(n.cfg.Socket))
	if errRes != nil {
		return errRes
	}

	listener, errLis := net.ListenTCP("tcp", tcpAddr)
	if errLis != nil {
		return errLis
	}

	go func() {
		fmt.Printf("listening on port %s\n", n.cfg.Socket)

		select {
		case sign := <-n.chShutdown:
			{
				fmt.Printf("\nshutting node on received: %s.\n", sign)
				close(n.chShutdown)
			}

		default:
			{
				rpc.Accept(listener)
			}
		}
	}()

	return nil
}
