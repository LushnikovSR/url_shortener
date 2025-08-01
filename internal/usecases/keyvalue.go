package usecases

import "errors"

type KeyValueRepository interface {
	Set(key, value string)
	Get(key string) (string, bool)
}

type KeyValueUsecase struct {
	repo KeyValueRepository
}

func NewKeyValueUsecase(repo KeyValueRepository) *KeyValueUsecase {
	return &KeyValueUsecase{repo: repo}
}

func (u *KeyValueUsecase) Add(key, value string) error {
	_, exists := u.repo.Get(key)
	if exists {
		return errors.New("key exists")
	}
	u.repo.Set(key, value)
	return nil
}

func (u *KeyValueUsecase) Get(key string) (string, error) {
	value, exists := u.repo.Get(key)
	if !exists {
		return "", errors.New("key not found")
	}
	return value, nil
}
