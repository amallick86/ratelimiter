package ratelimiter

import "sync"

type Limiter struct {
	buckets map[string]*Bucket
	mutex   sync.Mutex
	config  Config
}

var _ RateLimiter = (*Limiter)(nil) // compile-time type checking vi

func NewLimiter(config Config) *Limiter {
	return &Limiter{
		buckets: make(map[string]*Bucket),
		config:  config,
	}
}

func (l *Limiter) getBucket(ip string) *Bucket {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	if bucket, exists := l.buckets[ip]; exists {
		return bucket
	}
	bucket := NewBucket(l.config.Capacity, l.config.RefillRate)
	l.buckets[ip] = bucket
	return bucket
}

func (l *Limiter) Allow(ip string) bool {
	bucket := l.getBucket(ip)
	return bucket.Allow()
}
