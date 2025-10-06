package memory

import (
	"errors"
	"sync"
	"unicode/utf8"
)

type SafeMap struct {
	mu   sync.RWMutex
	data map[string]string
}

func New(initialSize int) *SafeMap {
	return &SafeMap{
		data: make(map[string]string, initialSize),
	}
}

func (sm *SafeMap) Set(key, value string) error {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	if utf8.RuneCountInString(key) < 4 {
		return errors.New("key must have four characters or more")
	}
	sm.data[key] = value
	return nil
}

func (sm *SafeMap) Get(key string) (string, bool) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	val, ok := sm.data[key]
	return val, ok
}
