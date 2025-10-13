package rainbowcolor

import (
	"github.com/roydevashish/golang-design-patterns/adapter/pkg/rainbow"
	"github.com/roydevashish/golang-design-patterns/adapter/video"
)

type RainbowColor struct {
	rainbow *rainbow.Rainbow
}

func NewRainbowColor(r *rainbow.Rainbow) *RainbowColor {
	return &RainbowColor{
		rainbow: r,
	}
}

func (r *RainbowColor) Apply(video *video.Video) {
	r.rainbow.Setup()
	r.rainbow.Update(video)
}
