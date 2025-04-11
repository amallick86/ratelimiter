package ratelimiter_test

import (
	"testing"

	"github.com/amallick86/ratelimiter"
)

func TestLimiter_Allow(t *testing.T) {
	t.Run("should limit requests per IP", func(t *testing.T) {
		config := ratelimiter.Config{
			Capacity:   3,
			RefillRate: 1,
		}
		limiter := ratelimiter.NewLimiter(config)
		
		ip1 := "192.168.1.1"
		ip2 := "192.168.1.2"
		
		// First 3 requests from IP1 should be allowed
		for i := 0; i < 3; i++ {
			if !limiter.Allow(ip1) {
				t.Errorf("Request %d from IP1 should be allowed", i+1)
			}
		}
		
		// 4th request from IP1 should be denied
		if limiter.Allow(ip1) {
			t.Error("4th request from IP1 should be denied")
		}
		
		// Requests from IP2 should still be allowed (different bucket)
		for i := 0; i < 3; i++ {
			if !limiter.Allow(ip2) {
				t.Errorf("Request %d from IP2 should be allowed", i+1)
			}
		}
		
		// 4th request from IP2 should be denied
		if limiter.Allow(ip2) {
			t.Error("4th request from IP2 should be denied")
		}
	})
}