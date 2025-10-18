package clouddata

import "fmt"

type CloudData struct {
	url string
}

func NewCloudData(url string) *CloudData {
	return &CloudData{
		url: url,
	}
}

func (c *CloudData) Save(data string) {
	fmt.Println("saving data", data, "to cloud at url:", c.url)
}
