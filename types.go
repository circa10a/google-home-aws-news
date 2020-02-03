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
	Items []CarouselItem `json:"items"`
}

type SimpleResponse struct {
	TextToSpeech string `json:"textToSpeech"`
}

type Items struct {
	SimpleResponse SimpleResponse `json:"simpleResponse"`
	CarouselBrowse CarouselBrowse `json:"carouselBrowse"`
}

type RichResponse struct {
	Items []Items `json:"items"`
}

type Google struct {
	ExpectUserResponse bool         `json:"expectUserResponse"`
	RichResponse       RichResponse `json:"richResponse"`
}

type Payload struct {
	Google Google `json:"google"`
}
