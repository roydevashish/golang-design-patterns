package encryptiondecorator

import (
	"encoding/base64"
	"fmt"

	"github.com/roydevashish/golang-design-patterns/decorator/data"
	datadecorator "github.com/roydevashish/golang-design-patterns/decorator/data_decorator"
)

type EncryptionDecorator struct {
	datadecorator.DataDecorator
}

func NewEncryptionDecorator(data data.Data) *EncryptionDecorator {
	return &EncryptionDecorator{
		*datadecorator.NewDataDecorator(data),
	}
}

func (e *EncryptionDecorator) Save(data string) {
	fmt.Println("encrypting data")
	e.DataDecorator.Save(e.Encrypt(data))
}

func (e *EncryptionDecorator) Encrypt(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}
