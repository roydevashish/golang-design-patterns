package handler

import "github.com/roydevashish/golang-design-patterns/chain-of-responsibility/http"

type Handler interface {
	SetNext(next Handler) Handler
	Handle(req http.HTTPRequest)
}
