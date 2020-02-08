package main

import (
	"time"

	awsnews "github.com/circa10a/go-aws-news/news"
	"github.com/patrickmn/go-cache"
	log "github.com/sirupsen/logrus"
)

func createCache(d time.Duration, c time.Duration) *cache.Cache {
	return cache.New(d, c)
}

func getNewsFromCache() awsnews.Announcements {
	news, found := Cache.Get(CacheKey)
	if found {
		return news.(awsnews.Announcements)
	}
	setNewsInCache()
	return getNewsFromCache()
}

func setNewsInCache() {
	news, err := awsnews.FetchYear(time.Now().Year())
	log.Info("News fetched")
	if err != nil {
		log.Error(err)
	}
	Cache.Set(CacheKey, news.Last(10), cache.DefaultExpiration)
	log.Info("Cache renewed")
}
