package router

import (
	"net/http"
	"urlshortener/app"

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

func (router *Router) redirect(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortUrl := vars["shorturl"]

	fullUrl, err := router.app.DB.GetFullUrlFromShortUrl(shortUrl)
	if err != nil {
		router.app.Logger.Info("No Url found for request with shorturl \"%s\"", shortUrl)
	}
	http.Redirect(w, r, fullUrl, http.StatusSeeOther)
}

func (router *Router) InitRouter(app *app.App) {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/{shorturl}", router.redirect).Methods("GET")

	app.Logger.Panic(http.ListenAndServe(":8080", myRouter).Error())
}
