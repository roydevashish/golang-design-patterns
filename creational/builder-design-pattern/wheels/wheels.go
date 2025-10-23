package wheels

type Wheels struct {
	diameterInInches float32
}

func NewWheels(diameterInInches float32) *Wheels {
	return &Wheels{
		diameterInInches: diameterInInches,
	}
}

func (w *Wheels) GetDiameterInInches() float32 {
	return w.diameterInInches
}
