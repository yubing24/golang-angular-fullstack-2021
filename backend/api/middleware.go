package api

import "net/http"

func ResponseInterceptor(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PUT")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		h.ServeHTTP(w, r)
	}
}