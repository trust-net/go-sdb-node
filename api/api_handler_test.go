package api

import (
    "testing"
    "bytes"
    "strings"
    "net/http"
    "net/http/httptest"
)

func getBody(w *httptest.ResponseRecorder) string {
	buf := new(bytes.Buffer)
    buf.ReadFrom(w.Result().Body)
    return buf.String()	
}

func TestGetError(t *testing.T) {
	handler := NewHandler(func(r *http.Request) (ApiResponse, Error) {
			return nil, ApiError(ERR_NOT_FOUND, "not found")
	})
	r, w := httptest.NewRequest("GET", "http://example.com", nil), httptest.NewRecorder()
	handler.ServeHTTP(w, r)
	body := getBody(w)
	if w.Result().StatusCode != ERR_NOT_FOUND || !strings.Contains(body, "not found") {
		t.Errorf("Unexpected Status: %d, Body: %s", w.Result().StatusCode, body)
	}
}

func TestGetWithBody(t *testing.T) {
	handler := NewHandler(func(r *http.Request) (ApiResponse, Error) {
			return "test response", nil
	})
	r, w := httptest.NewRequest("GET", "http://example.com", nil), httptest.NewRecorder()
	handler.ServeHTTP(w, r)
	body := getBody(w)
	if w.Result().StatusCode != 200 || !strings.Contains(body, "test response") {
		t.Errorf("Unexpected Status: %d, Body: %s", w.Result().StatusCode, body)
	}
}


func TestGetWithoutBody(t *testing.T) {
	handler := NewHandler(func(r *http.Request) (ApiResponse, Error) {
			return nil, nil
	})
	r, w := httptest.NewRequest("GET", "http://example.com", nil), httptest.NewRecorder()
	handler.ServeHTTP(w, r)
	body := getBody(w)
	if w.Result().StatusCode != 200 || len(body) > 0 {
		t.Errorf("Unexpected Status: %d, Body: %s", w.Result().StatusCode, body)
	}
}

func TestPostWithBody(t *testing.T) {
	handler := NewHandler(func(r *http.Request) (ApiResponse, Error) {
			return "test response", nil
	})
	r, w := httptest.NewRequest("POST", "http://example.com", nil), httptest.NewRecorder()
	handler.ServeHTTP(w, r)
	body := getBody(w)
	if w.Result().StatusCode != 201 || !strings.Contains(body, "test response") {
		t.Errorf("Unexpected Status: %d, Body: %s", w.Result().StatusCode, body)
	}
}

func TestPostWithoutBody(t *testing.T) {
	handler := NewHandler(func(r *http.Request) (ApiResponse, Error) {
			return nil, nil
	})
	r, w := httptest.NewRequest("POST", "http://example.com", nil), httptest.NewRecorder()
	handler.ServeHTTP(w, r)
	body := getBody(w)
	if w.Result().StatusCode != 202 || len(body) > 0 {
		t.Errorf("Unexpected Status: %d, Body: %s", w.Result().StatusCode, body)
	}
}

func TestPutWithBody(t *testing.T) {
	handler := NewHandler(func(r *http.Request) (ApiResponse, Error) {
			return "test response", nil
	})
	r, w := httptest.NewRequest("PUT", "http://example.com", nil), httptest.NewRecorder()
	handler.ServeHTTP(w, r)
	body := getBody(w)
	if w.Result().StatusCode != 202 || !strings.Contains(body, "test response") {
		t.Errorf("Unexpected Status: %d, Body: %s", w.Result().StatusCode, body)
	}
}

func TestPutWithoutBody(t *testing.T) {
	handler := NewHandler(func(r *http.Request) (ApiResponse, Error) {
			return nil, nil
	})
	r, w := httptest.NewRequest("PUT", "http://example.com", nil), httptest.NewRecorder()
	handler.ServeHTTP(w, r)
	body := getBody(w)
	if w.Result().StatusCode != 202 || len(body) > 0 {
		t.Errorf("Unexpected Status: %d, Body: %s", w.Result().StatusCode, body)
	}
}

func TestDeleteWithBody(t *testing.T) {
	handler := NewHandler(func(r *http.Request) (ApiResponse, Error) {
			return "test response", nil
	})
	r, w := httptest.NewRequest("DELETE", "http://example.com", nil), httptest.NewRecorder()
	handler.ServeHTTP(w, r)
	body := getBody(w)
	if w.Result().StatusCode != 200 || !strings.Contains(body, "test response") {
		t.Errorf("Unexpected Status: %d, Body: %s", w.Result().StatusCode, body)
	}
}

func TestDeleteWithoutBody(t *testing.T) {
	handler := NewHandler(func(r *http.Request) (ApiResponse, Error) {
			return nil, nil
	})
	r, w := httptest.NewRequest("DELETE", "http://example.com", nil), httptest.NewRecorder()
	handler.ServeHTTP(w, r)
	body := getBody(w)
	if w.Result().StatusCode != 204 || len(body) > 0 {
		t.Errorf("Unexpected Status: %d, Body: %s", w.Result().StatusCode, body)
	}
}

func TestUnknownMethod(t *testing.T) {
	handler := NewHandler(func(r *http.Request) (ApiResponse, Error) {
			return nil, nil
	})
	r, w := httptest.NewRequest("OPTION", "http://example.com", nil), httptest.NewRecorder()
	handler.ServeHTTP(w, r)
	body := getBody(w)
	if w.Result().StatusCode != ERR_METHOD_NOT_ALLOWED || !strings.Contains(body, "method not allowed") {
		t.Errorf("Unexpected Status: %d, Body: %s", w.Result().StatusCode, body)
	}
}
