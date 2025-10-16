package crop

type CropIconFactory struct {
	icons map[CropType]*CropIcon
}

func NewCropIconFactory() *CropIconFactory {
	return &CropIconFactory{
		icons: make(map[CropType]*CropIcon),
	}
}

func (c *CropIconFactory) GetCropIcon(cropType CropType) *CropIcon {
	_, ok := c.icons[cropType]
	if !ok {
		c.icons[cropType] = NewCropIcon(cropType, nil)
	}

	return c.icons[cropType]
}
