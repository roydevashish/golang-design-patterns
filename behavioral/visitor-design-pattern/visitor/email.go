package visitor

import "fmt"

type EmailVisitor struct{}

func (e *EmailVisitor) VisitLaw(l *LawClient) {
	fmt.Println("sending mail to", l.Email)
}

func (e *EmailVisitor) VisitRestaurent(r *RestaurentClient) {
	fmt.Println("sending mail to", r.Email)
}

func (e *EmailVisitor) VisitRetail(r *RetailClient) {
	fmt.Println("sending mail to", r.Email)
}
