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

	api := api.New(app)
	auth := auth.New(app)
	_ = auth
	myRouter.HandleFunc("/{shorturl}", api.Redirect).Methods("GET")

	app.Logger.Panic(http.ListenAndServe(":8080", myRouter).Error())
}
