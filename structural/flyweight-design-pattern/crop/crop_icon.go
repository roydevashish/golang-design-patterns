package crop

type CropIcon struct {
	cropType CropType
	icon     []byte
}

func NewCropIcon(cropType CropType, icon []byte) *CropIcon {
	return &CropIcon{
		cropType: cropType,
		icon:     icon,
	}
}

func (c *CropIcon) GetType() CropType {
	return c.cropType
}
