package compressormp4

import "fmt"

type CompressorMP4 struct{}

func NewCompressorMP4() *CompressorMP4 {
	return &CompressorMP4{}
}

func (c *CompressorMP4) Compress() {
	fmt.Println("compressing using MP4 compressor")
}
