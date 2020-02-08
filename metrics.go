package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	ginprometheus "github.com/zsais/go-gin-prometheus"
)

// Aggregating function to call all metric go routines
func recordMetrics() {
	countCacheItems()
}

func promMiddleware(g *gin.Engine) {
	p := ginprometheus.NewPrometheus("gin")
	p.Use(g)
}

func countCacheItems() {
	go func() {
		for {
			itemsInCache.Set(float64(len(getNewsFromCache())))
			// Update number of items in the cache every 30m
			// 10 healthy
			// 0 bad
			time.Sleep(30 * time.Minute)
		}
	}()
}

var (
	itemsInCache = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "google_cache_items",
		Help: "The total number of news items in the cache",
	})
)
