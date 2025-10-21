package mvc

type TwigController struct{}

func NewTwigController() *TwigController {
	return &TwigController{}
}

func (t *TwigController) CreateViewEngine() ViewEngine {
	return NewTwigViewEngine()
}
