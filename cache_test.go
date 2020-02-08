package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNewsFromCache(t *testing.T) {
	assert.Equal(t, len(getNewsFromCache()), 10)
}

func TestSetNewsInCache(t *testing.T) {
	setNewsInCache()
	assert.Equal(t, len(getNewsFromCache()), 10)
}
