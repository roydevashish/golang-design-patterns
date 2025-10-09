package visitor

type Client interface {
	Accept(v Visitor)
}

type Visitor interface {
	VisitLaw(l *LawClient)
	VisitRestaurent(r *RestaurentClient)
	VisitRetail(r *RetailClient)
}
