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

	apiPrefix := "/api/v1/"

	router.Handle(apiPrefix, api.auth.KeyAuth(api.AddNewShort)).Methods("POST")
	router.Handle(apiPrefix+"{shorturl}", api.auth.KeyAuth(api.GetFullFromShort)).Methods("GET")

	router.HandleFunc("/{shorturl}", api.Redirect).Methods("GET")

}
