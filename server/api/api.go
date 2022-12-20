package api

import (
	"net/http"
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

	router.Handle(apiPrefix, corsMiddleware(api.auth.KeyAuth(api.AddNewShort))).Methods("POST", "OPTIONS")
	router.Handle(apiPrefix, corsMiddleware(api.auth.KeyAuth(api.GetUrlsForAccessToken))).Methods("GET", "OPTIONS")
	router.Handle(apiPrefix+"{shorturl}", corsMiddleware(api.auth.KeyAuth(api.GetFullFromShort))).Methods("GET", "OPTIONS")

	router.HandleFunc("/{shorturl}", api.Redirect).Methods("GET", "OPTIONS")
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("content-type", "application/json;charset=UTF-8")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}
