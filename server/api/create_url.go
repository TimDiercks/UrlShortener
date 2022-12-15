package api

import (
	"encoding/json"
	"net/http"
	"urlshortener/auth"
	"urlshortener/db"
)

type RequestUrlParams struct {
	FullUrl  string `json:"url"`
	ShortUrl string `json:"myShort"`
}

type ResponseRequestUrl struct {
	ShortUrl string `json:"shortUrl"`
}

func (api *Api) RequestUrl(w http.ResponseWriter, r *http.Request) {
	var req RequestUrlParams

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		api.app.Logger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

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

	var shortUrl string
	if req.ShortUrl != "" && user.Role == "admin" {
		shortUrl = req.ShortUrl
	} else {
		shortUrl = db.GenerateShortUrl(apikey, req.FullUrl)
	}

	url, err := api.app.DB.InsertShortUrl(db.CreateShortUrlParams{
		UserId:   user.Id,
		ShortUrl: shortUrl,
		FullUrl:  req.FullUrl,
	})

	res, err := json.Marshal(ResponseRequestUrl{ShortUrl: url.ShortUrl})
	if err != nil {
		api.app.Logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(res)
}
