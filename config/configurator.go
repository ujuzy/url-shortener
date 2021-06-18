package config

import (
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
)

type Configuration struct {
	Server   ServerConfiguration
	Database DatabaseConfiguration
}

type ServerConfiguration struct {
	Port int
}

type DatabaseConfiguration struct {
	Address  string
	Port     int
	Basename string
	User     string
	Password string
}

func ReadConfig() *Configuration {
	err := godotenv.Load("./config/dev.env")
	if err != nil {
		log.Error("Error loading config file")
	}

	value, ok := os.LookupEnv("PORT")
	port, err := strconv.Atoi(value)
	if !ok || err != nil {
		log.Warn("No port passed. Using default 3001 port to run server.")
		port = 3001
	}

	value, ok = os.LookupEnv("PG_PORT")
	pgPort, err := strconv.Atoi(value)
	if !ok || err != nil {
		log.Warn("No postgres port passed. Using default 5432 PostgreSQL port")
		pgPort = 5432
	}

	return &Configuration{
		Server: ServerConfiguration{
			Port: port,
		},
		Database: DatabaseConfiguration{
			Address:  os.Getenv("PG_IP"),
			Port:     pgPort,
			Basename: os.Getenv("PG_DATABASE"),
			User:     os.Getenv("PG_USER"),
			Password: os.Getenv("PG_PASSWORD"),
		},
	}
}
