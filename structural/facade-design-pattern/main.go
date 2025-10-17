package main

import (
	orderrequest "github.com/roydevashish/golang-design-patterns/facade/order_request"
	orderservice "github.com/roydevashish/golang-design-patterns/facade/order_service"
)

func main() {
	orderRequest := orderrequest.NewOrderRequest()
	orderService := orderservice.NewOrderService()
	orderService.Order(orderRequest)
}
