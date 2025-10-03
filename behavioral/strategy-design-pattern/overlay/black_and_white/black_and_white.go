package blackandwhite

import "fmt"

type BlackAndWhite struct{}

func NewBlackAndWhite() *BlackAndWhite {
	return &BlackAndWhite{}
}

func (b *BlackAndWhite) Apply() {
	fmt.Println("applying black and white overlay")
}
