package main

import (
	"fmt"
	"go-url-shortener/handler"
	"go-url-shortener/store"

	"github.com/gin-gonic/gin"
)

func main() {

	// Note that store initialization happens here
	store.InitializeStore()

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hey Go URL Shortener !",
		})
	})

	r.GET("/favicon.ico", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"message": "favicon.ico",
		})
	})

	r.POST("/create-short-url", func(c *gin.Context) {
		handler.CreateShortUrl(c)
	})

	r.POST("/create-qrcode", func(c *gin.Context) {
		handler.CreateQRCode(c)
	})

	r.GET("/get-qrcode", func(c *gin.Context) {
		handler.GetQRCode(c)
	})

	r.GET("/:shortUrl", func(c *gin.Context) {
		handler.HandleShortUrlRedirect(c)
	})

	err := r.Run(":9808")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}
