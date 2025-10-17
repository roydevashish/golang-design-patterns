package orderservice

import (
	"github.com/roydevashish/golang-design-patterns/facade/authentication"
	"github.com/roydevashish/golang-design-patterns/facade/inventory"
	orderfulfillment "github.com/roydevashish/golang-design-patterns/facade/order_fulfillment"
	orderrequest "github.com/roydevashish/golang-design-patterns/facade/order_request"
	"github.com/roydevashish/golang-design-patterns/facade/payment"
)

type OrderService struct{}

func NewOrderService() *OrderService {
	return &OrderService{}
}

func (o *OrderService) Order(orderRequest *orderrequest.OrderRequest) {

	auth := authentication.NewAuthentication()
	auth.CheckAuth()

	inventory := inventory.NewInventory()
	for _, id := range orderRequest.ItemIds {
		inventory.CheckInventory(id)
	}

	payment := payment.NewPayment(orderRequest.Name, orderRequest.CardNumber, orderRequest.Amount)
	payment.Pay()

	orderFulfillment := orderfulfillment.NewOrderFulfillment(inventory)
	orderFulfillment.FulFill(orderRequest.Name, orderRequest.Address, orderRequest.ItemIds)
}
