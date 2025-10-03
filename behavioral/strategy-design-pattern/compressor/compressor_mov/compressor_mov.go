package compressormov

import "fmt"

type CompressorMOV struct{}

func NewCompressorMOV() *CompressorMOV {
	return &CompressorMOV{}
}

func (c *CompressorMOV) Compress() {
	fmt.Println("compressing using MOV compressor")
}
