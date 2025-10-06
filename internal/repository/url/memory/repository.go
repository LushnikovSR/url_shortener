package memory

import (
	db "github.com/LushnikovSR/url_shortener/pkg/memory"
)

type Repository struct {
	store *db.SafeMap
}

func NewRepository(instance *db.SafeMap) *Repository {
	return &Repository{store: instance}
}

func (r *Repository) Set(key, value string) error {
	err := r.store.Set(key, value)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) Get(key string) (string, bool) {
	return r.store.Get(key)
}
