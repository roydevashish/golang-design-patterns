package compressiondecorator

import (
	"encoding/base64"
	"fmt"

	"github.com/roydevashish/golang-design-patterns/decorator/data"
	datadecorator "github.com/roydevashish/golang-design-patterns/decorator/data_decorator"
)

type CompressionDecorator struct {
	datadecorator.DataDecorator
}

func NewCompressionDecorator(data data.Data) *CompressionDecorator {
	return &CompressionDecorator{
		*datadecorator.NewDataDecorator(data),
	}
}

func (c *CompressionDecorator) Save(data string) {
	fmt.Println("compressing data")
	c.DataDecorator.Save(c.Compress(data))
}

func (c *CompressionDecorator) Compress(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}
