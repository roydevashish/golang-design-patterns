package color

import "github.com/roydevashish/golang-design-patterns/adapter/video"

type Color interface {
	Apply(video *video.Video)
}
