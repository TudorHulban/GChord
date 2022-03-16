package fingers

import (
	"fmt"
	"net"
	"net/rpc"
)

type RPCServer struct {
	logic   interface{}
	chStop  chan struct{}
	chReady chan struct{}

	port string
}

func (s *RPCServer) Start() error {
	tcpAddr, errRes := net.ResolveTCPAddr("tcp", s.port)
	if errRes != nil {
		return fmt.Errorf("resolve TCP Addr: %w", errRes)
	}

	listener, errLis := net.ListenTCP("tcp", tcpAddr)
	if errLis != nil {
		return fmt.Errorf("listen TCP: %w", errLis)
	}

	rpc.Register(s.logic)

	fmt.Printf("listening on port %s\n", s.port)

	go rpc.Accept(listener)

	s.chReady <- struct{}{}

	for range s.chStop {
		listener.Close()
		fmt.Println("exiting loop...")
		break
	}

	fmt.Printf("stopped listening on port %s\n", s.port)

	return nil
}

func (s *RPCServer) GetReadyChannel() chan struct{} {
	return s.chReady
}

func (s *RPCServer) Stop() {
	s.chStop <- struct{}{}
}

func (s *RPCServer) CleanUp() {
	close(s.chReady)
	close(s.chStop)
}

func NewRPCServer(port string, logic interface{}) (*RPCServer, error) {
	return &RPCServer{
		logic:   logic,
		port:    port,
		chStop:  make(chan struct{}),
		chReady: make(chan struct{}),
	}, nil
}
