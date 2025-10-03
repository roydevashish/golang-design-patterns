package main

import (
	compressormov "github.com/roydevashish/golang-design-patterns/strategy/compressor/compressor_mov"
	compressormp4 "github.com/roydevashish/golang-design-patterns/strategy/compressor/compressor_mp4"
	"github.com/roydevashish/golang-design-patterns/strategy/overlay/blur"
	"github.com/roydevashish/golang-design-patterns/strategy/overlay/grey"
	videostorage "github.com/roydevashish/golang-design-patterns/strategy/video_storage"
)

func main() {
	vs := videostorage.NewVideoStorage(compressormp4.NewCompressorMP4(), blur.NewBlur())

	vs.SetCompressor(compressormov.NewCompressorMOV())

	vs.SetOverlay(grey.NewGrey())

	vs.Store("/save/final")
}
