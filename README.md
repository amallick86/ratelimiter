# ratelimiter


## `ratelimiter` — Token Bucket Rate Limiting for Go

`ratelimiter` is a lightweight, extensible rate-limiting package in Go based on the **Token Bucket algorithm**. It helps developers prevent abuse and manage traffic by limiting the number of allowed requests per IP or user over a given time window.

Ideal for APIs, microservices, and any system that needs robust request throttling.

---

### Features

- **Token Bucket Algorithm** — Refill-based limiter to handle burst and sustained traffic.
- **IP/User-Based Limiting** — Track and limit requests per individual client.
- **HTTP Middleware Support** — Easily integrate with `net/http` to guard your endpoints.
- **Configurable Limits** — Define bucket size and refill rate for full control.
- **Thread-Safe** — Designed with concurrency and high-load environments in mind.
- Lightweight & Fast — No external dependencies, perfect for production use.

---

### Use Cases

- API rate limiting per client IP.
- Prevent brute-force login attempts.
- Manage resource usage in multi-tenant systems.
- Throttle expensive backend operations.

---

### Configuration

```go
ratelimiter.Config{
    Capacity:   10, // Max number of requests allowed in a burst
    RefillRate: 2,  // Requests per second
}
```

---

### Install

```bash
go get github.com/amallick86/ratelimiter
```

---

### Example
 
 ```bash
 package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/amallick86/ratelimiter"
)

func main() {
	limiter := ratelimiter.NewLimiter(ratelimiter.Config{
		Capacity:   5,
		RefillRate: 1, // 1 token per second
	})

	mux := http.NewServeMux()
	mux.Handle("/", ratelimiter.Middleware(limiter, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, world!")
	})))

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", mux)
}
```