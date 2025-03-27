package cache

import (
	"errors"
	"sync"
	"time"
)

type memoryEntry struct {
	token     string
	expiresAt time.Time
}

type MemoryTokenStore struct {
	data map[uint]memoryEntry
	mu   sync.RWMutex
}

func NewMemoryTokenStore() RefreshTokenStore {
	return &MemoryTokenStore{
		data: make(map[uint]memoryEntry),
	}
}

func (m *MemoryTokenStore) Save(userID uint, token string, ttl time.Duration) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.data[userID] = memoryEntry{
		token:     token,
		expiresAt: time.Now().Add(ttl),
	}
	return nil
}

func (m *MemoryTokenStore) Find(userID uint) (string, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	entry, ok := m.data[userID]
	if !ok {
		return "", errors.New("token not found")
	}

	if time.Now().After(entry.expiresAt) {
		return "", errors.New("token expired")
	}
	return entry.token, nil
}

func (m *MemoryTokenStore) Delete(userID uint) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.data, userID)
	return nil
}
