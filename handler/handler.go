package handler

import (
	"encoding/json"
	"go-url-shortener/qrcode"
	"go-url-shortener/shortener"
	"go-url-shortener/store"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Request model definition
type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
	UserId  string `json:"user_id" binding:"required"`
}

func CreateShortUrl(c *gin.Context) {

	var creationRequest UrlCreationRequest
	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortUrl := shortener.GenerateShortLink(creationRequest.LongUrl, creationRequest.UserId)
	store.SaveUrlMapping(shortUrl, creationRequest.LongUrl, creationRequest.UserId)

	host := "http://localhost:9808/"
	c.JSON(200, gin.H{
		"message":   "short url created successfully",
		"short_url": host + shortUrl,
	})
}

func HandleShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	initialUrl := store.RetrieveInitialUrl(shortUrl)
	c.Redirect(302, initialUrl)
}

func CreateQRCode(c *gin.Context) {

	data, _ := c.GetRawData()
	var body map[string]string
	_ = json.Unmarshal(data, &body)
	url := body["url"]
	if url == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "url is required"})
		return
	}
	png := qrcode.GenerateQRCode(body["url"])
	c.Header("Content-Type", "image/png")
	c.String(200, string(png))

}

func GetQRCode(c *gin.Context) {

	url := c.Query("url")
	if url == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "url is required"})
		return
	}
	png := qrcode.GenerateQRCode(url)

	c.String(200, string(png))
}
