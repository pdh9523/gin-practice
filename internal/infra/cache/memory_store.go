package cache

import (
	"errors"
	"sync"
	"time"
)

type memoryEntry struct {
	value     string
	expiresAt time.Time
}

type MemoryTokenStore struct {
	data map[string]memoryEntry
	mu   sync.RWMutex
}

func NewMemoryTokenStore() GlobalCacheStore {
	return &MemoryTokenStore{
		data: make(map[string]memoryEntry),
	}
}

func (m *MemoryTokenStore) Save(key, value string, ttl time.Duration) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.data[key] = memoryEntry{
		value:     value,
		expiresAt: time.Now().Add(ttl),
	}
	return nil
}

func (m *MemoryTokenStore) Find(key string) (string, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	entry, ok := m.data[key]
	if !ok {
		return "", errors.New("cache not found")
	}

	if time.Now().After(entry.expiresAt) {
		return "", errors.New("cache expired")
	}
	return entry.value, nil
}

func (m *MemoryTokenStore) Delete(key string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.data, key)
	return nil
}
