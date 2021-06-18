package main

import (
	"url-shortener/config"
	"url-shortener/internal/app"
)

func main() {
	conf := config.ReadConfig()

	privateApp := app.New()
	privateApp.Run(conf.Server.Port)
}
