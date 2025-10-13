package blackandwhite

import (
	"fmt"

	"github.com/roydevashish/golang-design-patterns/adapter/video"
)

type BlackAndWhite struct{}

func (b *BlackAndWhite) Apply(video *video.Video) {
	fmt.Println("applying black and white color to the video.")
}
