package webserver

import (
	"github.com/roydevashish/golang-design-patterns/chain-of-responsibility/handler"
	"github.com/roydevashish/golang-design-patterns/chain-of-responsibility/http"
)

type WebServer struct {
	handler handler.Handler
}

func NewWebServer(handler handler.Handler) *WebServer {
	return &WebServer{
		handler: handler,
	}
}

func (w *WebServer) Handle(req http.HTTPRequest) {
	w.handler.Handle(req)
}
