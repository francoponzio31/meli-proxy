package app

import (
	"io"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

func ProxyController(c *gin.Context) {

	// Rate limit validation
	rateLimitMiddleware := RateLimiter("/proxy")
	rateLimitMiddleware(c)
	if c.IsAborted() {
		return
	}

	// Create a new HTTP request with the same method, headers and body
	config := LoadConfig()
	targetURL, _ := url.Parse(config.MeliAPIHost)
	targetURL.Path = c.Request.URL.Path
	targetURL.RawQuery = c.Request.URL.RawQuery

	req, err := http.NewRequest(c.Request.Method, targetURL.String(), c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error creating proxy request"})
		return
	}

	req.Header = c.Request.Header
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error doing doing Mercado Libre Request"})
		return
	}
	defer resp.Body.Close()

	// Copy request response to client request
	c.Status(resp.StatusCode)
	for k, v := range resp.Header {
		c.Header(k, v[0])
	}

	body, _ := io.ReadAll(resp.Body)
	c.Writer.Write(body)
}
