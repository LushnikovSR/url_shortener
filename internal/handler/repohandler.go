package handler

import (
	"net/http"
	"url_shortener/internal/usecase/dto"
)

type RepoHandler struct {
	repo Repository
}

func NewRepoHandler(repo Repository) *RepoHandler {
	return &RepoHandler{
		repo: repo,
	}
}

func (rh *RepoHandler) Add(w http.ResponseWriter, r *http.Request) {
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

	if err := rh.repo.Set(key, value); err != nil { //вызываем метод из repo.go через интерфейс contract.go в handler (не contract.go в usecase)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write([]byte("Data successfully added"))
}

func (rh *RepoHandler) Get(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("k")

	if key == "" {
		http.Error(w, "key parameter is required", http.StatusBadRequest)
		return
	}

	val, err := rh.repo.Get(key) //вызываем метод из repo.go через интерфейс contract.go в handler (не contract.go в usecase)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
	respondJSON(w, dto.Response{Message: val})
}
