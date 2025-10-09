package visitor

type RestaurentClient struct {
	Email string
}

func (r *RestaurentClient) Accept(v Visitor) {
	v.VisitRestaurent(r)
}
