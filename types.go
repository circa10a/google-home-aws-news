package main

type Response struct {
	Payload Payload `json:"payload"`
}

type OptionInfo struct {
	Key      string   `json:"key"`
	Synonyms []string `json:"synonyms"`
}

type Image struct {
	URL               string `json:"url"`
	AccessibilityText string `json:"accessibilityText"`
}

type ListItem struct {
	OptionInfo  OptionInfo `json:"optionInfo"`
	Description string     `json:"description"`
	Image       Image      `json:"image"`
	Title       string     `json:"title"`
}

type ListSelect struct {
	Title string     `json:"title"`
	Items []ListItem `json:"items"`
}

type Data struct {
	Type       string     `json:"@type"`
	ListSelect ListSelect `json:"listSelect"`
}

type SystemIntent struct {
	Intent string `json:"intent"`
	Data   Data   `json:"data"`
}

type SimpleResponse struct {
	TextToSpeech string `json:"textToSpeech"`
}

type Items struct {
	SimpleResponse SimpleResponse `json:"simpleResponse"`
}

type RichResponse struct {
	Items []Items `json:"items"`
}

type Google struct {
	ExpectUserResponse bool         `json:"expectUserResponse"`
	SystemIntent       SystemIntent `json:"systemIntent"`
	RichResponse       RichResponse `json:"richResponse"`
}

type Payload struct {
	Google Google `json:"google"`
}
