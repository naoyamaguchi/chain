package middleware

import "net/http"

func Example(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {}
	return http.HandlerFunc(fn)
}
