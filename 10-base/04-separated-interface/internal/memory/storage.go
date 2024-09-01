package memory

import (
	"errors"
	"sync"
)

// MemoryStorage implements the Storage interface using an in-memory map
type MemoryStorage struct {
	data map[string]string
	mu   sync.RWMutex
}

// NewMemoryStorage creates a new instance of MemoryStorage
func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		data: make(map[string]string),
	}
}

// Save stores the key-value pair in memory
func (s *MemoryStorage) Save(key string, value string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = value
	return nil
}

// Load retrieves the value associated with the key from memory
func (s *MemoryStorage) Load(key string) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	value, ok := s.data[key]
	if !ok {
		return "", errors.New("key not found")
	}
	return value, nil
}
