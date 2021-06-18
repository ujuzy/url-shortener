package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	uTranslator "url-shortener/internal/usecases/translator"
)

type App struct {
	translator *uTranslator.UrlTranslator
}

func New() *App {
	translator := uTranslator.New()

	return &App{
		translator: translator,
	}
}

func (h *App) Run(port int) {
	router := gin.Default()

	router.POST("/short", h.getShortUrl)
	router.POST("/long", h.getLongUrl)

	address := fmt.Sprintf("localhost:%d", port)
	err := router.Run(address)
	if err != nil {
		return
	}
}
