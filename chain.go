// Package alice provides a convenient way to chain http handlers.
package chain

import (
	"net/http"
)

type Middleware func(http.Handler) http.Handler

type Chains struct {
	Handler http.Handler
}

func NewChain(h http.Handler) *Chains {
	return &Chains{
		Handler: h,
	}
}

// func (c Chains) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// c.ServeHTTP(w, r)
// }

func (c *Chains) Chain(middleware Middleware) {
	c.Handler = middleware(c.Handler)
}
