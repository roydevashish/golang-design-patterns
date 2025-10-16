package crop

import "fmt"

type Crop struct {
	x        int
	y        int
	cropIcon *CropIcon
}

func NewCrop(x, y int, cropIcon *CropIcon) *Crop {
	return &Crop{
		x:        x,
		y:        y,
		cropIcon: cropIcon,
	}
}

func (c *Crop) Render() {
	fmt.Println("drawing", c.cropIcon.GetType(), "at", c.x, "and", c.y)
}
