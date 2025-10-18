package main

import (
	clouddata "github.com/roydevashish/golang-design-patterns/decorator/cloud_data"
	compressiondecorator "github.com/roydevashish/golang-design-patterns/decorator/compression_decorator"
	"github.com/roydevashish/golang-design-patterns/decorator/data"
	encryptiondecorator "github.com/roydevashish/golang-design-patterns/decorator/encryption_decorator"
)

func main() {
	url := "google.cloud.com"
	rawData := "this is some random data."
	encrypt := true
	compress := true

	var cloudData data.Data = clouddata.NewCloudData(url)

	if encrypt {
		cloudData = encryptiondecorator.NewEncryptionDecorator(cloudData)
	}

	if compress {
		cloudData = compressiondecorator.NewCompressionDecorator(cloudData)
	}

	cloudData.Save(rawData)
}
