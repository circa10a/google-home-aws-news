package main

import (
	"fmt"
	"net/http"

	awsnews "github.com/circa10a/go-aws-news/news"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func newsStatement(n []ListItem) string {
	if len(n) == 0 {
		return "No news from A.W.S yet."
	}
	return "Here's the latest A.W.S news"
}

func newsListItems() []ListItem {
	var newsItems []ListItem
	news, _ := awsnews.Fetch(2020, 01)
	for i, newsItem := range news {
		itemInfo := ListItem{
			Title:       newsItem.Title,
			Description: newsItem.PostDate,
			Image: Image{
				URL:               "https://a0.awsstatic.com/libra-css/images/logos/aws_logo_smile_1200x630.png",
				AccessibilityText: newsItem.Title,
			},
			OptionInfo: OptionInfo{
				Key:      fmt.Sprintf("NEWS_ITEM_%v", i),
				Synonyms: []string{},
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
				SystemIntent: SystemIntent{
					Intent: "actions.intent.OPTION",
					Data: Data{
						Type: "type.googleapis.com/google.actions.v2.OptionValueSpec",
						ListSelect: ListSelect{
							Title: "AWS News",
							Items: news,
						},
					},
				},
				RichResponse: RichResponse{
					Items: []Items{
						{
							SimpleResponse{
								TextToSpeech: newsStatement(news),
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
