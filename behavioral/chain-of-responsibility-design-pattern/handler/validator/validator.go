package validator

import (
	"fmt"

	"github.com/roydevashish/golang-design-patterns/chain-of-responsibility/handler"
	"github.com/roydevashish/golang-design-patterns/chain-of-responsibility/http"
)

type Validator struct {
	next handler.Handler
}

func NewValidator() *Validator {
	return &Validator{
		next: nil,
	}
}

func (v *Validator) SetNext(next handler.Handler) handler.Handler {
	v.next = next
	return next
}

func (v *Validator) Handle(req http.HTTPRequest) {
	fmt.Println("validating the request")

	if req.Username == "" || req.Password == "" {
		fmt.Println("validation failed")
		return
	}

	if v.next != nil {
		v.next.Handle(req)
	}
}
