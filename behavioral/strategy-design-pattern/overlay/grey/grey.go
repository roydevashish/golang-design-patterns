package grey

import "fmt"

type Grey struct{}

func NewGrey() *Grey {
	return &Grey{}
}

func (b *Grey) Apply() {
	fmt.Println("applying grey overlay")
}
