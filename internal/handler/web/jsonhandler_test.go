package web

import (
	"net/http/httptest"
	"testing"

	"github.com/LushnikovSR/url_shortener/internal/handler/dto"
)

func TestResponseJSON(t *testing.T) {
	w := httptest.NewRecorder()
	testResponse := dto.Response{Message: "test"}
	RespondJSON(w, testResponse)
	if w.Header().Get("Content-Type") != "application/json" {
		t.Error("Content-Type header is not application/json")
	}

	expectedBody := `{"message": "test"}`
	if w.Body.String() != expectedBody {
		t.Errorf("Expected body %s, got %s", expectedBody, w.Body.String())
	}
}
