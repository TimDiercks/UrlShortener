package auth

import (
	"errors"
	"net/http"
	"strings"
)

func (a *Auth) TokenAuth(endpoint func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if len(auth) == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		token, err := parseApiKey(auth)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		err = a.app.DB.VerifyToken(token)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		endpoint(w, r)
	})
}

func parseApiKey(header string) (string, error) {
	token := strings.Split(header, "Bearer ")
	if len(token) != 2 {
		return "", errTokenNotFound
	}

	return token[1], nil
}

var (
	errTokenNotFound = errors.New("could not find token in header")
)
