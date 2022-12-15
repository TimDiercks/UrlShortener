package api

import (
	"urlshortener/app"
	"urlshortener/auth"

	"github.com/gorilla/mux"
)

type Api struct {
	app  *app.App
	auth *auth.Auth
}

func New(app *app.App, auth *auth.Auth) *Api {
	return &Api{
		app:  app,
		auth: auth,
	}
}

func (api *Api) InitRoutes(router *mux.Router) {

	apiPrefix := "/v1/"

	router.Handle(apiPrefix+"get-shorturl", api.auth.KeyAuth(api.RequestUrl)).Methods("GET")

	router.HandleFunc("/{shorturl}", api.Redirect).Methods("GET")

}
