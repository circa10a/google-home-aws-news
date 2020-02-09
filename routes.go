package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleWebhook(c *gin.Context) {
	c.JSON(http.StatusOK, fulfillment())
}

// Init all context paths
func routes(router *gin.Engine) {
	root := router.Group("/")
	{
		root.POST("/webhook", handleWebhook)
	}
}
