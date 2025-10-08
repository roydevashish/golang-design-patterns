package logger

import (
	"fmt"

	"github.com/roydevashish/golang-design-patterns/chain-of-responsibility/handler"
	"github.com/roydevashish/golang-design-patterns/chain-of-responsibility/http"
)

type Logger struct {
	next handler.Handler
}

func NewLogger() *Logger {
	return &Logger{
		next: nil,
	}
}

func (l *Logger) SetNext(next handler.Handler) handler.Handler {
	l.next = next
	return next
}

func (l *Logger) Handle(req http.HTTPRequest) {
	fmt.Println("logging the request")

	fmt.Println(req.Username, req.Password)

	if l.next != nil {
		l.next.Handle(req)
	}
}
