package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"time"

	"github.com/naoyamaguchi/chain"
)

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("[START] middleware1")
		next.ServeHTTP(w, r)
		fmt.Println("[END] middleware1")
	})
}

type handler func(http.Handler) http.Handler

// Server is server struct
type Server struct {
	path []string
	// handler
}

func NewServer() (*Server, error) {
	server := &Server{
		path: []string{"^/hoge(.*)$", "^/piyo(.*)$", "^/guga(.*)$"},
	}
	return server, nil
}

func (server *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, p := range server.path {
		if regexp.MustCompile(p).Match([]byte(r.URL.String())) {
			fmt.Fprintln(w, "Hello from: "+r.URL.String())
			return
		}
	}
	// path not found
	http.NotFound(w, r)
}

func (server *Server) Run() {
	// chainedServer := chain.NewChain(middleware).Then(server)

	c := chain.NewChain(server)
	c.Chain(middleware)

	s := &http.Server{
		Addr: ":8080",
		// Handler: chainedServer,
		Handler:      c,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	if err := s.ListenAndServe(); err != nil {
		fmt.Println("[error] ", err)
	}
}

func main() {
	server, _ := NewServer()

	log.Println("Server start port 8080...")
	server.Run()

	// c := chain.NewChain(server)
	// c.Chain(middleware.Example)
	// c.Run()

}
