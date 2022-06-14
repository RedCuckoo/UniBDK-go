package api

import (
	"crypto/subtle"
	"fmt"
	"net/http"
)

func NewBasicAuthRpc(handler http.Handler, realm string, creds map[string]string) http.Handler {
	return &BasicAuthRPC{
		handler: handler,
		realm:   realm,
		creds:   creds,
	}
}

type BasicAuthRPC struct {
	handler http.Handler
	realm   string
	creds   map[string]string
}

func (ba *BasicAuthRPC) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user, pass, ok := r.BasicAuth()
	if !ok {
		basicAuthFailed(w, ba.realm)
		return
	}

	credPass, credUserOk := ba.creds[user]
	if !credUserOk || subtle.ConstantTimeCompare([]byte(pass), []byte(credPass)) != 1 {
		basicAuthFailed(w, ba.realm)
		return
	}

	ba.handler.ServeHTTP(w, r)
}

func basicAuthFailed(w http.ResponseWriter, realm string) {
	w.Header().Add("WWW-Authenticate", fmt.Sprintf(`Basic realm="%s"`, realm))
	w.WriteHeader(http.StatusUnauthorized)
}
