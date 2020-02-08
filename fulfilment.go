package main

type Response struct {
	Payload Payload `json:"payload"`
}

type Payload struct {
	Google Google `json:"google"`
}

type Google struct {
	ExpectUserResponse bool         `json:"expectUserResponse"`
	RichResponse       RichResponse `json:"richResponse,omitempty"`
}

type RichResponse struct {
	Items []Item `json:"items,omitempty"`
}

type Item struct {
	SimpleResponse *SimpleResponse `json:"simpleResponse,omitempty"`
	CarouselBrowse *CarouselBrowse `json:"carouselBrowse,omitempty"`
}

type SimpleResponse struct {
	TextToSpeech string `json:"textToSpeech"`
}

type CarouselBrowse struct {
	Items []CarouselItem `json:"items"`
}

type CarouselItem struct {
	Title         string        `json:"title"`
	OpenURLAction OpenURLAction `json:"openUrlAction"`
	Description   string        `json:"description,omitempty"`
}

type OpenURLAction struct {
	URL string `json:"url"`
}

// Build out the full struct for the JSON response
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
