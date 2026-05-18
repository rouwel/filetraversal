package server

import (
	"testing"
	"net/http"
	"net/http/httptest"
)

func TestHomeView(t *testing.T) {
	r := httptest.NewRequest("GET", "/", nil)
	q := httptest.NewRecorder()

	HomeView(q,r)
	if status := q.Code ; status != http.StatusOK {
		t.Errorf("not successful: %v , %v", status, http.StatusOK)
	}
}