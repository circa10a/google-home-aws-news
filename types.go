package main

type Response struct {
	Payload Payload `json:"payload"`
}

type OpenURLAction struct {
	URL string `json:"url"`
}

type Image struct {
	URL               string `json:"url"`
	AccessibilityText string `json:"accessibilityText"`
}

type CarouselItem struct {
	Title         string        `json:"title"`
	OpenURLAction OpenURLAction `json:"openUrlAction"`
	Description   string        `json:"description"`
	Footer        string        `json:"footer"`
	Image         Image         `json:"image"`
}

type CarouselBrowse struct {
	Items []CarouselItem `json:"items,omitempty"`
}

type SimpleResponse struct {
	TextToSpeech string `json:"textToSpeech,omitempty"`
}

type Item struct {
	SimpleResponse SimpleResponse `json:"simpleResponse,omitempty"`
	CarouselBrowse CarouselBrowse `json:"carouselBrowse,omitempty"`
}

type RichResponse struct {
	Items []Item `json:"items,omitempty"`
}

type Google struct {
	ExpectUserResponse bool         `json:"expectUserResponse"`
	RichResponse       RichResponse `json:"richResponse,omitempty"`
}

type Payload struct {
	Google Google `json:"google"`
}
