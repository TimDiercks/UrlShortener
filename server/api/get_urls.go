package api

import (
	"encoding/json"
	"net/http"
	"time"
	"urlshortener/auth"

	"github.com/gorilla/mux"
)

type FullUrlResponse struct {
	FullUrl string `json:"fullUrl"`
}

func (api *Api) GetFullFromShort(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortUrl := vars["shorturl"]

	fullUrl, err := api.app.DB.GetFullUrlFromShortUrl(shortUrl)
	if err != nil {
		api.app.Logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(FullUrlResponse{FullUrl: fullUrl})
	if err != nil {
		api.app.Logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(res)
}

type ShortUrlResponse struct {
	ShortUrl  string
	FullUrl   string
	CreatedAt time.Time
}

func (api *Api) GetUrlsForAccessToken(w http.ResponseWriter, r *http.Request) {
	apikey, err := auth.ApiKeyFromRequest(r)
	if err != nil {
		api.app.Logger.Error(err.Error())
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	user, err := api.app.DB.GetUserByApiKey(apikey)
	if err != nil {
		api.app.Logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	shortUrls, err := api.app.DB.GetShortUrlsByUserId(user.Id)
	if err != nil {
		api.app.Logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var resList []ShortUrlResponse

	for _, shortUrl := range shortUrls {
		resList = append(resList, ShortUrlResponse{
			ShortUrl:  shortUrl.ShortUrl,
			FullUrl:   shortUrl.FullUrl,
			CreatedAt: shortUrl.CreatedAt,
		})
	}

	res, err := json.Marshal(resList)
	if err != nil {
		api.app.Logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(res)

}
