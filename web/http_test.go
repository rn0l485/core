package web

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
	"net/http"
)

func TestHttpServer(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, client")
	})
	s := HttpServer(":8080", handler)

	assert.NotNil(t, s)
}
