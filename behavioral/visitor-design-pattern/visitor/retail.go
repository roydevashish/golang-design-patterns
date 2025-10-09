package visitor

type RetailClient struct {
	Email string
}

func (r *RetailClient) Accept(v Visitor) {
	v.VisitRetail(r)
}
