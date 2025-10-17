package payment

import "fmt"

type Payment struct {
	name       string
	cardNumber string
	amount     float32
}

func NewPayment(name string, cardNumber string, amount float32) *Payment {
	return &Payment{
		name:       name,
		cardNumber: cardNumber,
		amount:     amount,
	}
}

func (p *Payment) Pay() {
	fmt.Println("charging card with name", p.name)
}
