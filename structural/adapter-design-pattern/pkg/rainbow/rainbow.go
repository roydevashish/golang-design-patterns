package rainbow

import (
	"fmt"

	"github.com/roydevashish/golang-design-patterns/adapter/video"
)

type Rainbow struct{}

func NewRainbow() *Rainbow {
	return &Rainbow{}
}

func (r *Rainbow) Setup() {
	fmt.Println("setting up rainbow filter")
}

func (r *Rainbow) Update(video *video.Video) {
	fmt.Println("applying rainbow filter to the video")
}
