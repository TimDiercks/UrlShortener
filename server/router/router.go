package router

import (
	"net/http"
	"urlshortener/api"
	"urlshortener/app"
	"urlshortener/auth"

	"github.com/gorilla/mux"
)

type Router struct {
	app *app.App
}

func New(app *app.App) *Router {
	return &Router{
		app: app,
	}
}

func (router *Router) InitRouter(app *app.App) {
	myRouter := mux.NewRouter().StrictSlash(true)

	auth := auth.New(app)
	api := api.New(app, auth)

	api.InitRoutes(myRouter)

	app.Logger.Panic(http.ListenAndServe(":8080", myRouter).Error())
}
