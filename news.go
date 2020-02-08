package main

// Ensure the user is notified in the event there is no news/aws site is down
func defaultNewsStatement(n []CarouselItem) string {
	if len(n) == 0 {
		return "No cloud computing news yet."
	}
	return "Here's the latest cloud computing news."
}

// There needs to be at minimum to items in the carousel or user gets an error
func defaultNewsItem() []CarouselItem {
	var newsItems []CarouselItem
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

// Get news items from cache and generate list of structs for the JSON response
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
