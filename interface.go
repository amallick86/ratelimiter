package ratelimiter

type RateLimiter interface {
	Allow(ip string) bool
}
