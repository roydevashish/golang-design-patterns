package midnightpurple

import (
	"fmt"

	"github.com/roydevashish/golang-design-patterns/adapter/video"
)

type MidnightPurple struct{}

func (m *MidnightPurple) Apply(video *video.Video) {
	fmt.Println("applying midnight purple color to the video.")
}
