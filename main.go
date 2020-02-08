package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	log "github.com/sirupsen/logrus"
)

var (
	// Cache Global news cache
	Cache *cache.Cache
	// CacheKey references news items in the cache
	CacheKey string
	// DefaultExpiration expires cache after 8 hours
	DefaultExpiration = 8 * time.Hour
	// CleanupInterval purges expired items
	CleanupInterval = 8 * time.Hour
)

func init() {
	// Create cache on startup
	Cache = createCache(DefaultExpiration, CleanupInterval)
	log.Info("Cache created")
}

func main() {
	// Err placeholder
	var err error
	// Create gin engine
	r := gin.Default()
	// Add prometheus middleware
	promMiddleware(r)
	// Start custom metrics counters
	recordMetrics()
	// Initialize paths and handlers in routes.go
	routes(r)
	// Handle the errrrrrrs
	if err = r.Run(); err != nil {
		log.WithError(err).Fatal("Couldn't start server")
	}
}
