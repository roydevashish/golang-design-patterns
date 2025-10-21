package mvc

type BladeController struct{}

func NewBladeController() *BladeController {
	return &BladeController{}
}

func (b *BladeController) CreateViewEngine() ViewEngine {
	return NewBladeViewEngine()
}
