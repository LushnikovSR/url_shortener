package usecase

import (
	"errors"
)

type Repo struct {
	repository Repository
}

func NewRepo(instance Repository) *Repo {
	return &Repo{repository: instance}
}

func (r *Repo) Set(key, value string) error {
	_, exists := r.repository.Get(key) //вызываем метод из repositry.go через интерфейс contract.go в usecase (не contract.go в handler)
	if exists {
		return errors.New("key exists")
	}
	r.repository.Set(key, value) //вызываем метод из repositry.go через интерфейс contract.go в usecase (не contract.go в handler)
	return nil
}

func (r *Repo) Get(key string) (string, error) {
	value, exists := r.repository.Get(key) //вызываем метод из repositry.go через интерфейс contract.go в usecase (не contract.go в handler)
	if !exists {
		return "", errors.New("key not found")
	}
	return value, nil
}
