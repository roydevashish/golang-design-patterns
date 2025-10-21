package mvc

type BladeViewEngine struct{}

func NewBladeViewEngine() *BladeViewEngine {
	return &BladeViewEngine{}
}

func (b *BladeViewEngine) Render(filename string, data map[string]string) string {
	return "view render from " + filename + " by blade"
}
