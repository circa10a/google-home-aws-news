package main

import (
	"net/http"
	"time"

	awsnews "github.com/circa10a/go-aws-news/news"
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	log "github.com/sirupsen/logrus"
)

func defaultNewsStatement(n []CarouselItem) string {
	if len(n) == 0 {
		return "No cloud computing news yet."
	}
	return "Here's the latest cloud computing news."
}

func defaultNewsItem() []CarouselItem {
	var newsItems []CarouselItem
	// Needs at minimium 2 items
	for i := 1; i <= 2; i++ {
		newsItems = append(newsItems, CarouselItem{
			Title:       "No recent news",
			Description: "Check back soon",
			OpenURLAction: OpenURLAction{
				URL: "https://aws.amazon.com/new",
			},
		})
	}
	return newsItems
}

func getNewsListItems() []CarouselItem {
	newsItems := make([]CarouselItem, 0)
	news := getNewsFromCache()
	if len(news) == 0 {
		return defaultNewsItem()
	}

	for _, newsItem := range news {
		itemInfo := CarouselItem{
			Title:       newsItem.Title,
			Description: newsItem.PostDate,
			OpenURLAction: OpenURLAction{
				URL: newsItem.Link,
			},
		}
		newsItems = append(newsItems, itemInfo)
	}
	return newsItems
}

func fulfillment() *Response {
	news := getNewsListItems()

	return &Response{
		Payload{
			Google{
				ExpectUserResponse: false,
				RichResponse: RichResponse{
					Items: []Item{
						{
							SimpleResponse: &SimpleResponse{
								TextToSpeech: defaultNewsStatement(news),
							},
						},
						{
							CarouselBrowse: &CarouselBrowse{
								Items: news,
							},
						},
					},
				},
			},
		},
	}
}

func handleWebhook(c *gin.Context) {
	c.JSON(http.StatusOK, fulfillment())
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
	Cache = cache.New(DefaultExpiration, CleanupInterval)
	log.Info("Cache created")
}

func main() {
	var err error

	r := gin.Default()
	r.POST("/webhook", handleWebhook)

	if err = r.Run(); err != nil {
		log.WithError(err).Fatal("Couldn't start server")
	}
}
