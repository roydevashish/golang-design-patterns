package videostorage

import (
	"fmt"

	"github.com/roydevashish/golang-design-patterns/strategy/compressor"
	"github.com/roydevashish/golang-design-patterns/strategy/overlay"
)

type VideoStorage struct {
	compressor compressor.Compressor
	overlay    overlay.Overlay
}

func NewVideoStorage(compressor compressor.Compressor, overlay overlay.Overlay) *VideoStorage {
	return &VideoStorage{
		compressor: compressor,
		overlay:    overlay,
	}
}

func (v *VideoStorage) SetCompressor(compressor compressor.Compressor) {
	v.compressor = compressor
}

func (v *VideoStorage) SetOverlay(overlay overlay.Overlay) {
	v.overlay = overlay
}

func (v *VideoStorage) Store(f string) {
	v.compressor.Compress()
	v.overlay.Apply()
	fmt.Println("video file stored at", f)
}
