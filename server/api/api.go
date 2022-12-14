package api

import (
	"urlshortener/app"
)

type Api struct {
	app *app.App
}

func New(app *app.App) *Api {
	return &Api{
		app: app,
	}
}
