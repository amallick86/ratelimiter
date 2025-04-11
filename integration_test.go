package ratelimiter_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/amallick86/ratelimiter"
)

func TestRateLimiterIntegration(t *testing.T) {
	// Create a real limiter with a small capacity
	config := ratelimiter.Config{
		Capacity:   2,
		RefillRate: 1,
	}
	limiter := ratelimiter.NewLimiter(config)
	
	// Create a simple handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	
	// Apply middleware
	middleware := ratelimiter.Middleware(limiter, handler)
	
	// Test with the same IP
	ip := "192.168.1.100:8080"
	
	// First two requests should be allowed
	for i := 0; i < 2; i++ {
		req, _ := http.NewRequest("GET", "/", nil)
		req.RemoteAddr = ip
		w := httptest.NewRecorder()
		middleware.ServeHTTP(w, req)
		
		if w.Code != http.StatusOK {
			t.Errorf("Request %d should be allowed, got status %d", i+1, w.Code)
		}
	}
	
	// Third request should be denied
	req, _ := http.NewRequest("GET", "/", nil)
	req.RemoteAddr = ip
	w := httptest.NewRecorder()
	middleware.ServeHTTP(w, req)
	
	if w.Code != http.StatusTooManyRequests {
		t.Errorf("Expected status code %d, got %d", http.StatusTooManyRequests, w.Code)
	}
}