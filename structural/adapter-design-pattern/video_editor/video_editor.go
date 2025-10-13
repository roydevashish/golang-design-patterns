package videoeditor

import (
	"github.com/roydevashish/golang-design-patterns/adapter/color"
	"github.com/roydevashish/golang-design-patterns/adapter/video"
)

type VideoEditor struct {
	video *video.Video
}

func NewVideoEditor(video *video.Video) *VideoEditor {
	return &VideoEditor{
		video: video,
	}
}

func (v *VideoEditor) ApplyColor(color color.Color) {
	color.Apply(v.video)
}
