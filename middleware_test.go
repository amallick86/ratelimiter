package ratelimiter_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/amallick86/ratelimiter"
	"github.com/amallick86/ratelimiter/mocks"
)

func TestMiddleWare(t *testing.T) {
	t.Run("should allow request when limiter allows", func(t *testing.T) {
		mockLimiter := new(mocks.RateLimiter)
		mockLimiter.On("Allow", mock.Anything).Return(true)

		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})
		middleware := ratelimiter.Middleware(mockLimiter, handler)
		req, _ := http.NewRequest("GET", "/", nil)
		req.RemoteAddr = "192.168.1.100:8080"
		w := httptest.NewRecorder()
		middleware.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
		}
		mockLimiter.AssertCalled(t, "Allow", "192.168.1.100")
	})
	
	t.Run("should reject request when limiter denies", func(t *testing.T) {
		mockLimiter := new(mocks.RateLimiter)
		mockLimiter.On("Allow", mock.Anything).Return(false)

		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})
		middleware := ratelimiter.Middleware(mockLimiter, handler)
		req, _ := http.NewRequest("GET", "/", nil)
		req.RemoteAddr = "192.168.1.100:8080"
		w := httptest.NewRecorder()
		middleware.ServeHTTP(w, req)

		if w.Code != http.StatusTooManyRequests {
			t.Errorf("Expected status code %d, got %d", http.StatusTooManyRequests, w.Code)
		}
	})
	
	t.Run("should extract IP from X-Forwarded-For header", func(t *testing.T) {
		mockLimiter := new(mocks.RateLimiter)
		mockLimiter.On("Allow", "10.0.0.1").Return(true)

		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})
		middleware := ratelimiter.Middleware(mockLimiter, handler)
		req, _ := http.NewRequest("GET", "/", nil)
		req.RemoteAddr = "192.168.1.100:8080"
		req.Header.Set("X-Forwarded-For", "10.0.0.1")
		w := httptest.NewRecorder()
		middleware.ServeHTTP(w, req)

		mockLimiter.AssertCalled(t, "Allow", "10.0.0.1")
	})
}
