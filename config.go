package ratelimiter

type Config struct {
	Capacity   int // max number of requests per second
	RefillRate int // number of tokens added per second
}
