package main

import (
	"net/http"

	awsnews "github.com/circa10a/go-aws-news/news"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func newsStatement(n []CarouselItem) string {
	if len(n) == 0 {
		return "No news from A.W.S yet."
	}
	return "Here's the latest A.W.S news."
}

func newsListItems() []CarouselItem {
	var newsItems []CarouselItem
	news, _ := awsnews.Fetch(2020, 01)
	for _, newsItem := range news[:10] {
		itemInfo := CarouselItem{
			Title:       newsItem.Title,
			Description: newsItem.PostDate,
			OpenURLAction: OpenURLAction{
				URL: newsItem.Link,
			},
			Footer: newsItem.PostDate,
			Image: Image{
				URL:               "https://a0.awsstatic.com/libra-css/images/logos/aws_logo_smile_1200x630.png",
				AccessibilityText: newsItem.Title,
			},
		}
		newsItems = append(newsItems, itemInfo)
	}
	return newsItems
}

func buildFulfillment() *Response {
	news := newsListItems()

	return &Response{
		Payload{
			Google{
				ExpectUserResponse: false,
				RichResponse: RichResponse{
					Items: []Items{
						{
							SimpleResponse{
								TextToSpeech: newsStatement(news),
							},
							CarouselBrowse{
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
	c.JSON(http.StatusOK, buildFulfillment())
}

func main() {
	var err error

	r := gin.Default()
	r.POST("/webhook", handleWebhook)

	if err = r.Run(); err != nil {
		log.WithError(err).Fatal("Couldn't start server")
	}
}
