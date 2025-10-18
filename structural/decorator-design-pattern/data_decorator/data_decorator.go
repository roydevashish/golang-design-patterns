package datadecorator

import "github.com/roydevashish/golang-design-patterns/decorator/data"

type DataDecorator struct {
	data.Data
}

func NewDataDecorator(data data.Data) *DataDecorator {
	return &DataDecorator{
		data,
	}
}
