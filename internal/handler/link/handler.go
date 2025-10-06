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
	key := r.URL.Query().Get("k")
	value := r.URL.Query().Get("v")

	if key == "" {
		http.Error(w, "key parameter is required", http.StatusBadRequest)
		return
	}
	if value == "" {
		http.Error(w, "value parameter is required", http.StatusBadRequest)
		return
	}

	if err := rh.repo.Set(key, value); err != nil { //вызываем метод из repo.go через интерфейс contract.go
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write([]byte("Data successfully added"))
}

func (rh *Handler) Get(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("k")

	if key == "" {
		http.Error(w, "key parameter is required", http.StatusBadRequest)
		return
	}

	val, err := rh.repo.Get(key) //вызываем метод из repo.go через интерфейс contract.go в handler (не contract.go в usecase)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
	web.RespondJSON(w, dto.Response{Message: val})
}
