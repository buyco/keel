package helper

import (
	"github.com/gorilla/mux"
	"net/http"
	"net/url"
)

// GetRouteVars gets variables from URL
func GetRouteVars(r *http.Request) map[string]string {
	return mux.Vars(r)
}

// GetQueryVars gets variables from query string
func GetQueryVars(r *http.Request) url.Values {
	return r.URL.Query()
}
