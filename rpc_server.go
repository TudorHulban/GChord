package fingers

import (
	"fmt"
	"net"
	"net/rpc"
	"time"
)

// TODO: move to goroutine rpc accept
func NewRPCServer(logic interface{}, sock string) error {
	rpc.Register(logic)

	tcpAddr, errRes := net.ResolveTCPAddr("tcp", sock)
	if errRes != nil {
		return fmt.Errorf("resolve TCP Addr: %w", errRes)
	}

	listener, errLis := net.ListenTCP("tcp", tcpAddr)
	if errLis != nil {
		return fmt.Errorf("listen TCP: %s", errRes)
	}

	fmt.Printf("listening on port %s\n", sock)

	ticker := time.NewTicker(300 * time.Millisecond)

	select {
	case <-ticker.C:
		{
			fmt.Println("RPC listening ended")
		}

	default:
		{
			fmt.Println("RPC listening started")
			rpc.Accept(listener)
		}
	}

	return nil
}
