package ratelimiter_test

import (
	"testing"
	"time"

	"github.com/amallick86/ratelimiter"
)

func TestBucket_Allow(t *testing.T) {
	t.Run("should allow requests when bucket has tokens", func(t *testing.T) {
		bucket := ratelimiter.NewBucket(5, 1)
		
		// First 5 requests should be allowed (initial capacity)
		for i := 0; i < 5; i++ {
			if !bucket.Allow() {
				t.Errorf("Request %d should be allowed", i+1)
			}
		}
		
		// 6th request should be denied
		if bucket.Allow() {
			t.Error("Request should be denied after bucket is empty")
		}
	})
	
	t.Run("should refill tokens over time", func(t *testing.T) {
		// Create a bucket with 2 capacity and 1 token refill per second
		bucket := ratelimiter.NewBucket(2, 1)
		
		// Use up all tokens
		bucket.Allow()
		bucket.Allow()
		
		// Next request should be denied
		if bucket.Allow() {
			t.Error("Request should be denied when bucket is empty")
		}
		
		// Wait for refill
		time.Sleep(1100 * time.Millisecond) // Wait a bit more than 1 second
		
		// Should have 1 token now
		if !bucket.Allow() {
			t.Error("Request should be allowed after refill")
		}
		
		// Should be empty again
		if bucket.Allow() {
			t.Error("Bucket should be empty after using refilled token")
		}
	})
}