package main

import "net/http"

func JsonifyResponse(handler http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		handler.ServeHTTP(w, r)
	})
}
