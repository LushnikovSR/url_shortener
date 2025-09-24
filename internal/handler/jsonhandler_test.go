package handler

import (
	"net/http/httptest"
	"testing"
	"url_shortener/internal/usecase/dto"
)

func TestResponseJSON(t *testing.T) {
	w := httptest.NewRecorder()
	testResponse := dto.Response{Message: "test"}
	respondJSON(w, testResponse)
	if w.Header().Get("Content-Type") != "application/json" {
		t.Error("Content-Type header is not application/json")
	}

	expectedBody := `{"message": "test"}`
	if w.Body.String() != expectedBody {
		t.Errorf("Expected body %s, got %s", expectedBody, w.Body.String())
	}
}
