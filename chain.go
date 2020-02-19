// Package alice provides a convenient way to chain http handlers.
package chain

import "net/http"

// type Constructor func(http.Handler) http.Handler

// type Chains struct {
// 	constructors []Constructor
// }

// func NewChain(constructors ...Constructor) Chains {
// 	return Chains{append(([]Constructor)(nil), constructors...)}
// }

// func (c Chains) Chain(constructors ...Constructor) {
// 	c.constructors = append(c.constructors, constructors...)
// }

// func (c Chains) Then(h http.Handler) http.Handler {
// 	if h == nil {
// 		h = http.DefaultServeMux
// 	}
// 	for i := range c.constructors {
// 		h = c.constructors[len(c.constructors)-1-i](h)
// 	}
// 	return h
// }

//////////////////////////////////////////////////////////
type Constructor func(http.Handler) http.Handler

type Chains struct {
	constructors []Constructor
	handler      http.Handler
}

func NewChain(h http.Handler) *Chains {
	return &Chains{handler: h}
}

func (c Chains) Chain(constructors ...Constructor) Chains {
	return Chains{append(([]Constructor)(nil), constructors...)}
	// return Chains{append(c.constructors, constructors...)}
}

func (c Chains) Then(h http.Handler) http.Handler {
	if h == nil {
		h = http.DefaultServeMux
	}
	for i := range c.constructors {
		h = c.constructors[len(c.constructors)-1-i](h)
	}
	return h
}
