package fingers

type finger struct {
	ID       []byte // ID hash of (n + 2^i) mod (2^m)
	PointsTo Node
}

type fingerTable []*finger // TODO: move to map
