package main

import orderscontroller "github.com/roydevashish/golang-design-patterns/factory-method/orders_controller"

func main() {
	ocBlade := orderscontroller.NewOrdersControllerBlade()
	ocTwig := orderscontroller.NewOrdersControllerTwig()
	ocBlade.ListOrders()
	ocTwig.GetOrder(10)
}
