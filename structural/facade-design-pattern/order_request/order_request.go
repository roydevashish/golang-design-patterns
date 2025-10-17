package orderrequest

type OrderRequest struct {
	Name       string
	CardNumber string
	Amount     float32
	Address    string
	ItemIds    []string
}

func NewOrderRequest() *OrderRequest {
	return &OrderRequest{
		Name:       "John Doe",
		CardNumber: "1234 5678 0912",
		Amount:     20.99,
		Address:    "123, Springfield, Texas",
		ItemIds:    []string{"123", "423", "555", "989"},
	}
}
