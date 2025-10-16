package main

import (
	"fmt"

	"github.com/roydevashish/golang-design-patterns/flyweight/crop"
)

func main() {
	cs := crop.NewCropService(crop.NewCropIconFactory())
	for _, crop := range cs.GetCrops() {
		crop.Render()
		fmt.Println(crop)
	}
}
