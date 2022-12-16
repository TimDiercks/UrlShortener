package auth

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"net/http"
	"strings"
)

func ApiKeyFromRequest(r *http.Request) (string, error) {
	auth := r.Header.Get("Authorization")
	if len(auth) == 0 {
		return "", errTokenNotFound
	}

	token, err := parseApiKey(auth)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (a *Auth) KeyAuth(endpoint func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := ApiKeyFromRequest(r)
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

func (a *Auth) GenerateApiKey() string {
	b := make([]byte, apiKeyLength)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
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
	apiKeyLength     = 32
)
