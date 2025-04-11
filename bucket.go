package ratelimiter

import (
	"sync"
	"time"
)

type Bucket struct {
	capacity     int
	tokens       int
	refillRate   int //tokens per second
	lastRefilled time.Time
	mutex        sync.Mutex
}

func NewBucket(capacity, refillRate int) *Bucket {
	return &Bucket{
		capacity:     capacity,
		tokens:       capacity,
		refillRate:   refillRate,
		lastRefilled: time.Now(),
	}
}

func (b *Bucket) Allow() bool {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	now := time.Now()
	elapsed := now.Sub(b.lastRefilled).Seconds()      // time difference in seconds
	newtokens := int(elapsed * float64(b.refillRate)) // number of tokens to add
	if newtokens > 0 {
		b.tokens = min(b.tokens+newtokens, b.capacity) // adding minimum of newtokens and capacity
		b.lastRefilled = now
	}
	if b.tokens > 0 {
		b.tokens--
		return true
	}
	return false
}
