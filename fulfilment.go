package main

// Modeled after google docs
// https://developers.google.com/assistant/conversational/responses#browsing_carousel

// Response is the entire JSON payload response
type Response struct {
	Payload Payload `json:"payload"`
}

// Payload is a google defined higher structure, see https://developers.google.com/assistant/conversational/responses#browsing_carousel
type Payload struct {
	Google Google `json:"google"`
}

// Google is another google defined higher structure, see https://developers.google.com/assistant/conversational/responses#browsing_carousel
type Google struct {
	ExpectUserResponse bool         `json:"expectUserResponse"`
	RichResponse       RichResponse `json:"richResponse,omitempty"`
}

// RichResponse gives UI representation of the news data fetched
type RichResponse struct {
	Items []Item `json:"items,omitempty"`
}

// Item provides different interactions
type Item struct {
	SimpleResponse *SimpleResponse `json:"simpleResponse,omitempty"`
	CarouselBrowse *CarouselBrowse `json:"carouselBrowse,omitempty"`
}

// Simple response provides audio only feedback
type SimpleResponse struct {
	TextToSpeech string `json:"textToSpeech"`
}

// CarouselBrowse provides a UI list of items with hyperlinks
type CarouselBrowse struct {
	Items []CarouselItem `json:"items"`
}

// CarouselItem is a UI entry for each news item
type CarouselItem struct {
	Title         string        `json:"title"`
	OpenURLAction OpenURLAction `json:"openUrlAction"`
	Description   string        `json:"description,omitempty"`
}

// OpenURLAction provides a url to be opened in the client's browser when touched
type OpenURLAction struct {
	URL string `json:"url"`
}

// fulfillment builds out the full struct for the JSON response
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
