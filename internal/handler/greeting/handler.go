package greeting

import (
	"net/http"

	"github.com/LushnikovSR/url_shortener/internal/handler/dto"
	"github.com/LushnikovSR/url_shortener/internal/handler/web"
)

type GreetHandler struct {
	greeting greeting
}

func New(instance greeting) *GreetHandler {
	return &GreetHandler{
		greeting: instance,
	}
}

func (h *GreetHandler) Root(w http.ResponseWriter, r *http.Request) {
	msg := h.greeting.Hello()
	web.RespondJSON(w, dto.Response{Message: msg})
}

func (h *GreetHandler) Name(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("q")
	msg := h.greeting.HelloName(name)
	web.RespondJSON(w, dto.Response{Message: msg})
}
