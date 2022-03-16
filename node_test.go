package fingers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNodeCo(t *testing.T) {
	cfg := &Config{
		Socket: ":8000",
	}

	require.NoError(t, cfg.Validate())

	var o Operations

	n, errCo := NewNode(cfg, o)
	require.NoError(t, errCo, "constructor issues")
	require.NotNil(t, n)

	require.NoError(t, n.Start(), "start node")
}
