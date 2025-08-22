package url

import (
	db "url_shortener/pkg/db"
)

type Repository struct {
	store *db.SafeMap
}

func NewRepository(instance *db.SafeMap) *Repository {
	return &Repository{store: instance}
}

func (r *Repository) Set(key, value string) {
	r.store.Set(key, value)
}

func (r *Repository) Get(key string) (string, bool) {
	return r.store.Get(key)
}
