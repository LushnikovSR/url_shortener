//go:generate mockgen -source ${GOFILE} -package ${GOPACKAGE}_test -destination=mocks_test.go
package link

import (
	"database/sql"
	"errors"
)

type Link struct {
	repository repository
}

func New(instance repository) *Link {
	return &Link{repository: instance}
}

func (r *Link) Set(originalLink, shortLink string) error {
	_, err := r.repository.GetOriginalLink(shortLink) //вызываем метод из repositry.go через интерфейс contract.go в usecase (не contract.go в handler)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			r.repository.Create(originalLink, shortLink)
			return nil
		}
		return err
	}
	return errors.New("short link exists")
}

func (r *Link) Get(shortLink string) (string, error) {
	originalLink, err := r.repository.GetOriginalLink(shortLink) //вызываем метод из repositry.go через интерфейс contract.go в usecase (не contract.go в handler)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", errors.New("key not found")
		}
		return "", err
	}
	return originalLink, nil
}
