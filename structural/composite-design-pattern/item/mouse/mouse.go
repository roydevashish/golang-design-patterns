package mouse

type Mouse struct {
	price float32
}

func NewMouse(p float32) *Mouse {
	return &Mouse{
		price: p,
	}
}

func (m *Mouse) GetPrice() float32 {
	return m.price
}
