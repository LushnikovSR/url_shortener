package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

// MockRepository реализует интерфейс Repository для тестов
type MockRepository struct {
	SetFunc func(key, value string) error
	GetFunc func(key string) (string, error)
}

func (mr *MockRepository) Set(key, value string) error {
	return mr.SetFunc(key, value)
}

func (mr *MockRepository) Get(key string) (string, error) {
	return mr.GetFunc(key)
}

func TestRepoHandler_Add(t *testing.T) {
	tests := []struct {
		name           string
		queryParams    string
		mockSetFunc    func(string, string) error
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "successful add",
			queryParams:    "?k=test&v=value",
			mockSetFunc:    func(string, string) error { return nil },
			expectedStatus: http.StatusOK,
			expectedBody:   "Data successfully added",
		},
		{
			name:           "missing key parameter",
			queryParams:    "?v=value",
			mockSetFunc:    func(string, string) error { return nil },
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "key parameter is required\n",
		},
		{
			name:           "missing value parameter",
			queryParams:    "?k=test",
			mockSetFunc:    func(string, string) error { return nil },
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "value parameter is required\n",
		},
		{
			name:           "repository error",
			queryParams:    "?k=test&v=value",
			mockSetFunc:    func(string, string) error { return errors.New("storage error") },
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "storage error\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &MockRepository{
				SetFunc: tt.mockSetFunc,
			}
			handler := NewRepoHandler(mockRepo)
			req := httptest.NewRequest("GET", "/add"+tt.queryParams, nil)
			w := httptest.NewRecorder()
			handler.Add(w, req)
			resp := w.Result()
			if resp.StatusCode != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, resp.StatusCode)
			}
			body := w.Body.String()
			if body != tt.expectedBody {
				t.Errorf("Expected body %s, got %s", tt.expectedBody, body)
			}
		})
	}
}

func TestRepoHandler_Get(t *testing.T) {
	tests := []struct {
		name           string
		queryParams    string
		mockGetFunc    func(string) (string, error)
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "successful get",
			queryParams:    "?k=test&v=value",
			mockGetFunc:    func(string) (string, error) { return "test_value", nil },
			expectedStatus: http.StatusOK,
			expectedBody:   `{"message":"test_value"}`,
		},
		{
			name:           "missed key parameter",
			queryParams:    "?v=value",
			mockGetFunc:    func(string) (string, error) { return "", nil },
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "key parameter is required\n",
		},
		{
			name:           "not found",
			queryParams:    "?k=nonexistent",
			mockGetFunc:    func(string) (string, error) { return "", errors.New("not found") },
			expectedStatus: http.StatusNotFound,
			expectedBody:   "not found\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &MockRepository{
				GetFunc: tt.mockGetFunc,
			}
			handler := NewRepoHandler(mockRepo)
			req := httptest.NewRequest("GET", "/get"+tt.queryParams, nil)
			w := httptest.NewRecorder()
			handler.Get(w, req)
			resp := w.Result()
			if resp.StatusCode != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, resp.StatusCode)
			}
			body := w.Body.String()
			if body != tt.expectedBody {
				t.Errorf("Expected body %s, got %s", body, tt.expectedBody)
			}
		})
	}
}
