package authenticator

import (
	"fmt"

	"github.com/roydevashish/golang-design-patterns/chain-of-responsibility/handler"
	"github.com/roydevashish/golang-design-patterns/chain-of-responsibility/http"
)

type Authenticator struct {
	next handler.Handler
}

func NewAuthenticator() *Authenticator {
	return &Authenticator{
		next: nil,
	}
}

func (a *Authenticator) SetNext(next handler.Handler) handler.Handler {
	a.next = next
	return next
}

func (a *Authenticator) Handle(req http.HTTPRequest) {
	fmt.Println("authenticating the request")

	if req.Username != "admin" || req.Password != "admin@123" {
		fmt.Println("authentication failed")
		return
	}

	if a.next != nil {
		a.next.Handle(req)
	}
}
