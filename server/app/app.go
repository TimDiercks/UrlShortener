package app

import (
	"urlshortener/config"
	"urlshortener/log"
)

type App struct {
	Logger *log.Logger
	Config *config.Config
}
