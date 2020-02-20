package middleware

import (
	"fmt"
	"net/http"
)

func Example1(next http.Handler) http.Handler {
	// fn := func(w http.ResponseWriter, r *http.Request) {}
	// return http.HandlerFunc(fn)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("[START] middleware 1")
		next.ServeHTTP(w, r)
		fmt.Println("[END] middleware 1")
	})
}

func Example2(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("[START] middleware 2")
		next.ServeHTTP(w, r)
		fmt.Println("[END] middleware 2")
	})
}
