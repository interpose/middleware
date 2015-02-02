package middleware

import (
	"net/http"
)

// NoPanic permits recovery from panics which occur during a request, so that
// errors which occur during a request do not bring down the server.
// from http://nicolasmerouze.com/middlewares-golang-best-practices-examples/
func NoPanic(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				http.Error(w, http.StatusText(500), 500)
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
