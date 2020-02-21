package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"time"

	"github.com/naoyamaguchi/chain"
	"github.com/naoyamaguchi/chain/cmd/middleware"
)

// Server is server struct
type Server struct {
	path []string
}

func NewServer() (*Server, error) {
	server := &Server{
		path: []string{"^/hoge$", "^/piyo(.*)$", "^/guga(.*)$"},
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
	// Chain middleware
	c := chain.NewChain(server)
	c.Chain(middleware.Example1)
	c.Chain(middleware.Example2)

	s := &http.Server{
		Addr:         ":8080",
		Handler:      c.Handler,
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

}
