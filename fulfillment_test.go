package main

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFulfillment(t *testing.T) {
	payload := fulfillment()
	_, jsonErr := json.Marshal(payload)
	assert.NoError(t, jsonErr)
}
