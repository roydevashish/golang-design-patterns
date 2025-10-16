package crop

type CropService struct {
	cropIconFactory *CropIconFactory
}

func NewCropService(cif *CropIconFactory) *CropService {
	return &CropService{
		cropIconFactory: cif,
	}
}

func (c *CropService) GetCrops() []Crop {
	var cropList []Crop

	carrot1 := NewCrop(1, 4, c.cropIconFactory.GetCropIcon(CARROT))
	carrot2 := NewCrop(1, 5, c.cropIconFactory.GetCropIcon(CARROT))
	carrot3 := NewCrop(1, 6, c.cropIconFactory.GetCropIcon(CARROT))

	cropList = append(cropList, *carrot1)
	cropList = append(cropList, *carrot2)
	cropList = append(cropList, *carrot3)

	return cropList
}
