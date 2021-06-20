package app

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func (h *App) getShortUrl(c *gin.Context) {
	request, err := readRequest(c)
	if err != nil {
		c.String(http.StatusInternalServerError, "Internal server error")
		return
	}

	if url, ok := request["url"]; ok {
		shortUrl, err := h.translator.ShortenUrl(c.Request.Context(), url)
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
	request, err := readRequest(c)
	if err != nil {
		c.String(http.StatusInternalServerError, "Internal server error")
		return
	}

	if url, ok := request["url"]; ok {
		longUrl, err := h.translator.ExtendUrl(c.Request.Context(), url)
		if err != nil {
			c.String(http.StatusInternalServerError, "Internal server error")
			return
		}
		c.JSONP(http.StatusOK, gin.H{"url": longUrl})
		return
	}

	c.String(http.StatusBadRequest, "Bad request")
}

func readRequest(c *gin.Context) (map[string]string, error) {
	reqBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return nil, err
	}

	var reqMap map[string]string
	err = json.Unmarshal(reqBody, &reqMap)
	if err != nil {
		return nil, err
	}

	return reqMap, nil
}
