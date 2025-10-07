package link

import (
	"net/http"

	"github.com/LushnikovSR/url_shortener/internal/handler/dto"
	"github.com/LushnikovSR/url_shortener/internal/handler/web"
)

type Handler struct {
	repo repository
}

func New(repo repository) *Handler {
	return &Handler{
		repo: repo,
	}
}

func (rh *Handler) Add(w http.ResponseWriter, r *http.Request) {
	ordinaryLink := r.URL.Query().Get("o")
	shortLink := r.URL.Query().Get("s")

	if ordinaryLink == "" {
		http.Error(w, "ordinaryLink link is required", http.StatusBadRequest)
		return
	}
	if shortLink == "" {
		http.Error(w, "shortLink link is required", http.StatusBadRequest)
		return
	}

	if err := rh.repo.Set(ordinaryLink, shortLink); err != nil { //вызываем метод из repo.go через интерфейс contract.go
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write([]byte("Data successfully added"))
}

func (rh *Handler) Get(w http.ResponseWriter, r *http.Request) {
	shortLink := r.URL.Query().Get("s")

	if shortLink == "" {
		http.Error(w, "shortLink link is required", http.StatusBadRequest)
		return
	}

	val, err := rh.repo.Get(shortLink) //вызываем метод из repo.go через интерфейс contract.go в handler (не contract.go в usecase)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
	web.RespondJSON(w, dto.Response{Message: val})
}
