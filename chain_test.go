package chain

import (
	"net/http"
	"testing"
)

type TestServer struct {
	path string
}

func (testServer *TestServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}

func TestNewChain(t *testing.T) {
	testServer := &TestServer{
		path: "hoge",
	}
	c := NewChain(testServer)
	if c == nil {
		t.Fatal("Failed test")
	}
}
