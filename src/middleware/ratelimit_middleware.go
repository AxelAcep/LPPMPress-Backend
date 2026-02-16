package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
)

type client struct {
	hits     int
	lastSeen time.Time
}

var (
	mu      sync.Mutex
	clients = make(map[string]*client)
)

func LoginRateLimiter() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ip := c.RealIP()

			mu.Lock()
			v, exists := clients[ip]

			if !exists {
				clients[ip] = &client{hits: 1, lastSeen: time.Now()}
				mu.Unlock()
				return next(c)
			}

			if time.Since(v.lastSeen) > 1*time.Hour {
				v.hits = 1
				v.lastSeen = time.Now()
				mu.Unlock()
				return next(c)
			}

			v.hits++
			v.lastSeen = time.Now()

			if v.hits > 15 {
				mu.Unlock()
				return c.JSON(http.StatusTooManyRequests, map[string]string{
					"message": "Cooldown Login for 1 hour!",
				})
			}

			mu.Unlock()
			return next(c)
		}
	}
}
