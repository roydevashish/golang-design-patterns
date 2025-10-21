package orderscontroller

import "github.com/roydevashish/golang-design-patterns/factory-method/mvc"

type OrdersController struct {
	*mvc.Controller
}

func NewOrdersControllerBlade() *OrdersController {
	return &OrdersController{
		mvc.NewController(mvc.NewBladeController()),
	}
}

func NewOrdersControllerTwig() *OrdersController {
	return &OrdersController{
		mvc.NewController(mvc.NewTwigController()),
	}
}

func (o *OrdersController) ListOrders() {
	orders := map[string]string{
		"red socks":    "$12.98",
		"blue socks":   "$12.98",
		"pink t-shirt": "29.00",
	}

	o.Render("orders.blade.php", orders)
}

func (o *OrdersController) GetOrder(id int) {
	order := map[string]string{"red socks": "$12.98"}
	o.Render("order.php", order)
}
