package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (api *Api) Redirect(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortUrl := vars["shorturl"]

	fullUrl, err := api.app.DB.GetFullUrlFromShortUrl(shortUrl)
	if err != nil {
		api.app.Logger.Info("No Url found for request with shorturl \"%s\"", shortUrl)
	}
	http.Redirect(w, r, fullUrl, http.StatusSeeOther)
}
