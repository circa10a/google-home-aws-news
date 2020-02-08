package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultNewsStatement(t *testing.T) {
	singleCarouselItem := []CarouselItem{
		{
			Title:       "Test",
			Description: "test",
			OpenURLAction: OpenURLAction{
				URL: "https://example.com",
			},
		},
	}

	assert.Equal(t, defaultNewsStatement(singleCarouselItem), "Here's the latest cloud computing news.")
	assert.Equal(t, defaultNewsStatement([]CarouselItem{}), "No cloud computing news yet.")
}

func TestDefaultNewsItem(t *testing.T) {
	assert.Equal(t, len(defaultNewsItem()), 2)
}

func TestGetNewsItems(t *testing.T) {
	assert.Equal(t, len(getNewsListItems()), 10)
}
