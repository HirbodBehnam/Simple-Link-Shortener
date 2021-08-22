package api

import (
	"RabinLink/config"
	"crypto/subtle"
	"net/http"
)

// authorizeHeader checks if the sender of this request has the right token or not
func authorizeHeader(r *http.Request) bool {
	return authorizeToken(r.Header.Get("Token"))
}

// authorizeToken safely checks if the token is valid or not
func authorizeToken(token string) bool {
	return subtle.ConstantTimeCompare([]byte(token), []byte(config.Config.Token)) == 1
}
