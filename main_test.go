package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

func TestCoreRoutes(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	pingGet, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, pingGet)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())

	k := httptest.NewRecorder()
	tenantsGet, _ := http.NewRequest("GET", "/tenants", nil)
	router.ServeHTTP(k, tenantsGet)

	assert.Equal(t, 200, k.Code)

}
