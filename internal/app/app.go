package app

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"url-shortener/config"
	dbTranslator "url-shortener/internal/domain/postgres/translator"
	uTranslator "url-shortener/internal/usecases/translator"
)

type App struct {
	config     *config.Configuration
	translator *uTranslator.UrlTranslator
}

func New(conf *config.Configuration) *App {
	pgConnString := fmt.Sprintf(
		"port=%d host=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Database.Port,
		conf.Database.Address,
		conf.Database.User,
		conf.Database.Password,
		conf.Database.Basename)

	postgres, err := sql.Open("postgres", pgConnString)
	if err != nil {
		log.Fatal("Error opening database")
		return nil
	}

	translatorService := dbTranslator.New(postgres)
	translator := uTranslator.New(translatorService)

	return &App{
		config:     conf,
		translator: translator,
	}
}

func (h *App) Run() {
	router := gin.Default()

	router.POST("/short", h.getShortUrl)
	router.POST("/long", h.getLongUrl)

	address := fmt.Sprintf("localhost:%d", h.config.Server.Port)
	err := router.Run(address)
	if err != nil {
		return
	}
}
