package fingers

import (
	"fmt"
	"net"
	"net/rpc"
)

func NewRPCServer(logic interface{}, s Socket) error {
	rpc.Register(logic)

	tcpAddr, errRes := net.ResolveTCPAddr("tcp", string(s))
	if errRes != nil {
		return fmt.Errorf("resolve TCP Addr: %w", errRes)
	}

	listener, errLis := net.ListenTCP("tcp", tcpAddr)
	if errLis != nil {
		return fmt.Errorf("listen TCP: %s", errRes)
	}

	fmt.Printf("listening on port %s\n", string(s))

	for {
		rpc.Accept(listener)
	}
}
