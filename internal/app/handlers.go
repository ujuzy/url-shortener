package app

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func (h *App) getShortUrl(c *gin.Context) {
	reqBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.String(http.StatusInternalServerError, "Internal server error")
		return
	}

	var reqMap map[string]string
	err = json.Unmarshal(reqBody, &reqMap)
	if err != nil {
		c.String(http.StatusInternalServerError, "Internal server error")
		return
	}

	if url, ok := reqMap["url"]; ok {
		shortUrl, err := h.translator.ShortenUrl(url)
		if err != nil {
			c.String(http.StatusInternalServerError, "Internal server error")
			return
		}
		c.JSONP(http.StatusOK, gin.H{"url": shortUrl})
		return
	}

	c.String(http.StatusBadRequest, "Bad request")
}

func (h *App) getLongUrl(c *gin.Context) {
	c.JSONP(http.StatusOK, gin.H{"url": "long-url-here"})
}
