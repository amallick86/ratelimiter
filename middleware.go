package ratelimiter

import (
	"net/http"
	"strings"
)

func Middleware(l RateLimiter, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := extractIP(r)
		if !l.Allow(ip) {
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func extractIP(r *http.Request) string {
	ip := r.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip = r.RemoteAddr
	}
	// strip port
	if colon := strings.LastIndex(ip, ":"); colon != -1 {
		// Handle IPv6 addresses which are enclosed in square brackets
		if strings.HasPrefix(ip, "[") && strings.Contains(ip[:colon], "]") {
			// For IPv6, remove the brackets as well
			ip = strings.Trim(ip[:strings.LastIndex(ip, "]")+1], "[]")
		} else {
			// For IPv4, just remove the port part
			ip = ip[:colon]
		}
	}
	return ip
}
