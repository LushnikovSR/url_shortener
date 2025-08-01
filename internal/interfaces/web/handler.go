package web

import (
	"encoding/json"
	"net/http"
	"url_shortener/internal/entities"
	"url_shortener/internal/usecases"
)

type Handler struct {
	greetingUC *usecases.GreetingUsecase
	keyValueUC *usecases.KeyValueUsecase
}

func NewHandler(greetingUC *usecases.GreetingUsecase, keyValueUC *usecases.KeyValueUsecase) *Handler {
	return &Handler{
		greetingUC: greetingUC,
		keyValueUC: keyValueUC,
	}
}

func (h *Handler) Root(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, entities.Response{Message: h.greetingUC.Hello()})
}

func (h *Handler) Name(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("q")
	respondJSON(w, entities.Response{Message: h.greetingUC.HelloName(name)})
}

func (h *Handler) Add(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("k")
	value := r.URL.Query().Get("v")
	if err := h.keyValueUC.Add(key, value); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write([]byte("Data successfully added"))
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("k")
	val, err := h.keyValueUC.Get(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
	respondJSON(w, entities.Response{Message: val})
}

func respondJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
