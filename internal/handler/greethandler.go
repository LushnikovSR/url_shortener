package handler

import (
	"net/http"
	"url_shortener/internal/usecase/dto"
)

type GreetHandler struct {
	greeting Greeting
}

func NewGreetHandler(instance Greeting) *GreetHandler {
	return &GreetHandler{
		greeting: instance,
	}
}

func (h *GreetHandler) Root(w http.ResponseWriter, r *http.Request) {
	msg := h.greeting.Hello()
	respondJSON(w, dto.Response{Message: msg})
}

func (h *GreetHandler) Name(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("q")
	msg := h.greeting.HelloName(name)
	respondJSON(w, dto.Response{Message: msg})
}
