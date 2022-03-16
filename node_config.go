package fingers

import "time"

type Config struct {
	Hash     hasher
	HashSize int

	Socket string

	StabilizeMin time.Duration
	StabilizeMax time.Duration

	Timeout time.Duration
	MaxIdle time.Duration
}

func (c *Config) Validate() error {
	return nil
}
