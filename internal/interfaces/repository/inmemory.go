package repository

import (
	"url_shortener/pkg/safemap"
)

type InMemoryRepo struct {
	store *safemap.SafeMap
}

func NewInMemoryRepo(store *safemap.SafeMap) *InMemoryRepo {
	return &InMemoryRepo{store: store}
}

func (r *InMemoryRepo) Set(key, value string) {
	r.store.Set(key, value)
}

func (r *InMemoryRepo) Get(key string) (string, bool) {
	return r.store.Get(key)
}
