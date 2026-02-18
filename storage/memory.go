package storage

import (
	"errors"
	"sync"
	"url-shortener/model"
)

type MemoryStorage struct {
	mu   sync.RWMutex
	data map[string]*model.URL
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		data: make(map[string]*model.URL),
	}
}

func (m *MemoryStorage) Save(url *model.URL) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[url.Code] = url
}

func (m *MemoryStorage) Get(code string) (*model.URL, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	url, exists := m.data[code]
	if !exists {
		return nil, errors.New("not found")
	}
	return url, nil
}

func (m *MemoryStorage) IncrementClicks(code string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if url, ok := m.data[code]; ok {
		url.Clicks++
	}
}
