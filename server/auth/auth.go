package auth

import "urlshortener/app"

type Auth struct {
	app *app.App
}

func New(app *app.App) *Auth {
	return &Auth{
		app: app,
	}
}
