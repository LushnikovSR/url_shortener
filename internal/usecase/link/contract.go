package link

import (
	"github.com/LushnikovSR/url_shortener/internal/repository/url/db"
)

/*
	type repository interface {
		Set(key, value string) error
		Get(key string) (string, bool)
	}
*/
type repository interface {
	Create(originalLink, shortLink string) (int, error)
	GetOriginalLink(shortLink string) (string, error)
	GetAllLinks() ([]db.Link, error)
	Update(id int, newShortLink string) error
	Delete(id int) error
}
