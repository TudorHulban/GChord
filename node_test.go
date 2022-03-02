package fingers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNodeCo(t *testing.T) {
	cfg := &Config{
		Socket: "0.0.0.0:8000",
	}

	require.NoError(t, cfg.Validate())

	n, errCo := NewNode(cfg)
	require.NoError(t, errCo, "constructor issues")
	require.NotNil(t, n)

	var o Operations
	require.NoError(t, n.Start(&o), "start node")
}
