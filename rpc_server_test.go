package fingers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewRPCServ(t *testing.T) {
	sock := ":8000"
	var o Operations

	_, errNew := NewRPCServer(sock, &o)
	require.NoError(t, errNew)
}
