package db

import "sync"

type SafeMap struct {
	mu   sync.RWMutex
	data map[string]string
}

func New(initialSize int) *SafeMap {
	return &SafeMap{
		data: make(map[string]string, initialSize),
	}
}

func (sm *SafeMap) Set(key, value string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.data[key] = value
}

func (sm *SafeMap) Get(key string) (string, bool) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	val, ok := sm.data[key]
	return val, ok
}
