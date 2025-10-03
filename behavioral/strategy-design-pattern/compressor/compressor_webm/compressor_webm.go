package compressorwebm

import "fmt"

type CompressorWebm struct{}

func NewCompressorWebm() *CompressorWebm {
	return &CompressorWebm{}
}

func (c *CompressorWebm) Compress() {
	fmt.Println("compressing using Webm compressor")
}
