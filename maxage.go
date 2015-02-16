package middleware

import (
	"fmt"
	"net/http"
)

// MaxAgeHandler sets a cache header to indicate when a resource should expire.
// Taken from https://groups.google.com/forum/#!topic/golang-nuts/n-GjwsDlRco
func MaxAgeHandler(seconds int, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Cache-Control", fmt.Sprintf("max-age=%d, public, must-revalidate, proxy-revalidate", seconds))
		h.ServeHTTP(w, r)
	})
}
