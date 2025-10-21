package mvc

import "fmt"

type Controller struct {
	IController
}

func NewController(cont IController) *Controller {
	return &Controller{
		cont,
	}
}

func (c *Controller) Render(filename string, data map[string]string) {
	viewEngine := c.CreateViewEngine()
	html := viewEngine.Render(filename, data)
	fmt.Println(html)
}

type IController interface {
	CreateViewEngine() ViewEngine
}
