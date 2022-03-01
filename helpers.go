package fingers

import "hash"

type hasher func() hash.Hash

func hashWith(h hash.Hash, key string) []byte {
	h.Write([]byte(key))

	return h.Sum(nil)
}
