package main

import (
	"github.com/roydevashish/golang-design-patterns/chain-of-responsibility/handler/authenticator"
	"github.com/roydevashish/golang-design-patterns/chain-of-responsibility/handler/logger"
	"github.com/roydevashish/golang-design-patterns/chain-of-responsibility/handler/validator"
	"github.com/roydevashish/golang-design-patterns/chain-of-responsibility/http"
	webserver "github.com/roydevashish/golang-design-patterns/chain-of-responsibility/web_server"
)

func main() {
	val := validator.NewValidator()
	auth := authenticator.NewAuthenticator()
	log := logger.NewLogger()

	val.SetNext(auth).SetNext(log)

	serv := webserver.NewWebServer(val)

	req1 := http.NewHTTPRequest("admin", "admin@123")
	serv.Handle(*req1)
}
