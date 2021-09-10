package env

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPIGet(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/env", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(New().APIGet)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wront status code: got %v want %v", status, http.StatusOK)
	}
}
