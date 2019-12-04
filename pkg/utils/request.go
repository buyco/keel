package utils

import (
	"github.com/gorilla/mux"
	"net/http"
	"net/url"
)

func GetRouteVars(r *http.Request) map[string]string {
	return mux.Vars(r)
}

func GetQueryVars(r *http.Request) url.Values {
	return r.URL.Query()
}
