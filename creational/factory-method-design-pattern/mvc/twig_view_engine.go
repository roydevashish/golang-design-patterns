package mvc

type TwigViewEngine struct{}

func NewTwigViewEngine() *TwigViewEngine {
	return &TwigViewEngine{}
}

func (t *TwigViewEngine) Render(filename string, data map[string]string) string {
	return "view render from " + filename + " by twig"
}
