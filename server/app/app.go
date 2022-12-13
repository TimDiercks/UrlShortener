package app

import (
	"urlshortener/config"
	"urlshortener/db"
	"urlshortener/log"
)

type App struct {
	Logger *log.Logger
	Config *config.Config
	DB     *db.DB
}
