package main

import (
	blackandwhite "github.com/roydevashish/golang-design-patterns/adapter/color/black_and_white"
	midnightpurple "github.com/roydevashish/golang-design-patterns/adapter/color/midnight_purple"
	rainbowcolor "github.com/roydevashish/golang-design-patterns/adapter/color/rainbow_color"
	"github.com/roydevashish/golang-design-patterns/adapter/pkg/rainbow"
	"github.com/roydevashish/golang-design-patterns/adapter/video"
	videoeditor "github.com/roydevashish/golang-design-patterns/adapter/video_editor"
)

func main() {
	ve := videoeditor.NewVideoEditor(video.NewVideo())
	ve.ApplyColor(&blackandwhite.BlackAndWhite{})
	ve.ApplyColor(&midnightpurple.MidnightPurple{})
	ve.ApplyColor(rainbowcolor.NewRainbowColor(&rainbow.Rainbow{}))

}
