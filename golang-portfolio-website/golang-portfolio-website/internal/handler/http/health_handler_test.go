package http

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealth_ReturnsOK(t *testing.T) {
	h := NewHealthHandler()
	rr := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/health", nil)
	h.Health(rr, r)
	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rr.Code)
	}
}
