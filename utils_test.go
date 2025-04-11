package ratelimiter

import (
	"net/http"
	"testing"
)

func TestExtractIP(t *testing.T) {
	t.Run("should extract IP from RemoteAddr", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/", nil)
		req.RemoteAddr = "192.168.1.100:8080"
		
		ip := extractIP(req)
		
		if ip != "192.168.1.100" {
			t.Errorf("Expected IP 192.168.1.100, got %s", ip)
		}
	})
	
	t.Run("should extract IP from X-Forwarded-For header", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/", nil)
		req.RemoteAddr = "192.168.1.100:8080"
		req.Header.Set("X-Forwarded-For", "10.0.0.1")
		
		ip := extractIP(req)
		
		if ip != "10.0.0.1" {
			t.Errorf("Expected IP 10.0.0.1, got %s", ip)
		}
	})
	
	t.Run("should handle IPv6 addresses", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/", nil)
		req.RemoteAddr = "[2001:db8::1]:8080"
		
		ip := extractIP(req)
		
		if ip != "2001:db8::1" {
			t.Errorf("Expected IP 2001:db8::1, got %s", ip)
		}
	})
}