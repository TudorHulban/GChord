package fingers

import (
	"hash"
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

	chShutdown    chan struct{}
	lastStablized time.Time
}
