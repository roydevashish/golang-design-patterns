package visitor

type LawClient struct {
	Email string
}

func (l *LawClient) Accept(v Visitor) {
	v.VisitLaw(l)
}
