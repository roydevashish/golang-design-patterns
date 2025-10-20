package testing

import (
	"fmt"

	appsettings "github.com/roydevashish/golang-design-patterns/singleton/app_settings"
)

func Run() {
	as := appsettings.GetInstance()
	fmt.Println("from testing pkg:", as.Get("sub-title"))
}
