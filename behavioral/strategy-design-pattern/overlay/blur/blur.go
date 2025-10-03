package blur

import "fmt"

type Blur struct{}

func NewBlur() *Blur {
	return &Blur{}
}

func (b *Blur) Apply() {
	fmt.Println("applying blur overlay")
}
