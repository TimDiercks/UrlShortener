package main

import (
	"os"
	"urlshortener/app"
	"urlshortener/config"
	"urlshortener/db"
	"urlshortener/log"
	"urlshortener/router"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	var err error
	var app app.App

	// Logger
	app.Logger = log.New()
	app.Logger.Info("Urlshortener starting...")
	if len(os.Args) <= 1 {
		app.Logger.Panic("Config is missing")
	}

	// Config
	configPath := os.Args[1]
	app.Logger.Info("Reading config \"%s\"", configPath)
	app.Config, err = config.Parse(configPath)
	if err != nil {
		app.Logger.Panic("Could not parse config", configPath)
	}

	// Database
	app.Logger.Info("Connecting to database...")
	dbConn, err := sqlx.Open("postgres", app.Config.DB)
	if err != nil {
		app.Logger.Panic(err.Error())
	}
	if err = dbConn.Ping(); err != nil {
		app.Logger.Panic("Database connection could not be established")
	}
	app.DB = db.New(dbConn)
	app.Logger.Info("Database connection established")

	// Router
	router := router.New(&app)
	router.InitRouter(&app)
}
