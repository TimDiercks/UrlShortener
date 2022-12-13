package main

import (
	"database/sql"
	"os"
	"urlshortener/app"
	"urlshortener/config"
	"urlshortener/db"
	"urlshortener/log"

	_ "github.com/lib/pq"
)

func main() {
	var err error
	var app app.App

	// Logger
	app.Logger = log.New()
	app.Logger.Info("Urlshortener starting...")
	if len(os.Args) <= 1 {
		app.Logger.Error("Config is missing")
		return
	}

	// Config
	configPath := os.Args[1]
	app.Logger.Info("Reading config \"%s\"", configPath)
	app.Config, err = config.Parse(configPath)
	if err != nil {
		app.Logger.Error("Could not parse config", configPath)
		return
	}

	// Database
	app.Logger.Info("Connecting to database...")
	dbConn, err := sql.Open("postgres", app.Config.DB)
	if err != nil {
		app.Logger.Panic(err.Error())
	}
	app.DB = db.New(dbConn)
	app.Logger.Info("Database connection established")

}
