package middleware

import (
	"net/http"
)

func HandlerVaryOnHTMXRequest(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(
        func(w http.ResponseWriter, r *http.Request) {
            w.Header().Add("Vary", "HX-Request")
            next.ServeHTTP(w, r)
        },
    )
}
