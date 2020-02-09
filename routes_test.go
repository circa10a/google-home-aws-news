package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gavv/httpexpect"
	"github.com/stretchr/testify/assert"
	//"fmt"
)

func init() {
	routes(G)
	promMiddleware(G)
}

func TestRulesEndpoint(t *testing.T) {
	// run server using httptest
	server := httptest.NewServer(G)
	defer server.Close()

	// create httpexpect instance
	e := httpexpect.New(t, server.URL)
	// is it working?
	response := e.POST("/webhook").
		Expect().
		Status(http.StatusOK).JSON()
	// Ensure correct news items
	response.Object().Value("payload").Object().Value("google").Object().ValueEqual("expectUserResponse", false)
	response.Object().Value("payload").
		Object().Value("google").
		Object().Value("richResponse").
		Object().Value("items").Array().Element(0).
		Object().Value("simpleResponse").
		Object().ValueEqual("textToSpeech", "Here's the latest cloud computing news.")

	news := response.Object().Value("payload").
		Object().Value("google").
		Object().Value("richResponse").
		Object().Value("items").Array().Element(1).
		Object().Value("carouselBrowse").Object().Value("items").Array()

	// Ensure cache and payload line up
	newsCache := getNewsFromCache()
	assert.Equal(t, len(news.Iter()), len(newsCache))

	for i := range news.Iter() {
		news.Element(i).Object().ValueEqual("title", newsCache[i].Title)
		news.Element(i).Object().ValueEqual("description", newsCache[i].PostDate)
		news.Element(i).Object().Value("openUrlAction").Object().ValueEqual("url", newsCache[i].Link)
	}
}

func TestMetricsEndpoint(t *testing.T) {
	// run server using httptest
	server := httptest.NewServer(G)
	defer server.Close()

	// create httpexpect instance
	e := httpexpect.New(t, server.URL)

	// is it working?
	response := e.GET("/metrics").
		Expect().
		Status(http.StatusOK).ContentType("text/plain")
	// Ensure proper refirect to swagger docs
	response.Body().Contains("gin_request_duration_seconds_sum")
}
