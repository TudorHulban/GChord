package fingers

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNewRPCServ(t *testing.T) {
	sock := ":8000"
	var o Operations

	s, errNew := NewRPCServer(sock, &o)
	require.NoError(t, errNew)

	go func() {
		<-s.GetReadyChannel()

		time.Sleep(100 * time.Millisecond)
		s.Stop()
	}()

	require.NoError(t, s.Start())
}
