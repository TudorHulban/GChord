package fingers

import "sync"

type hashKey string

type store struct {
	data map[hashKey]string
	Hash hasher
	mu   sync.RWMutex
}

func (s *store) getKey(k hashKey) ([]byte, error) {
	s.mu.RLock()
	val, exists := s.data[k]
	s.mu.RUnlock()

	if !exists {
		return nil, errKeyNotFound
	}

	return []byte(val), nil
}

func (s *store) Set(k hashKey, value string) {
	s.mu.Lock()
	s.data[k] = value
	s.mu.Unlock()
}

func (s *store) Delete(k hashKey) {
	s.mu.Lock()
	delete(s.data, k)
	s.mu.Unlock()
}
