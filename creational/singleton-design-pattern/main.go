package main

import (
	"fmt"

	appsettings "github.com/roydevashish/golang-design-patterns/singleton/app_settings"
	"github.com/roydevashish/golang-design-patterns/singleton/testing"
)

func main() {
	as := appsettings.GetInstance()
	as.Set("title", "design pattern")
	as.Set("sub-title", "with go lang")
	fmt.Println("from main func:", as.Get("sub-title"))

	testing.Run()
}
