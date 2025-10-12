package keyboard

type KeyBoard struct {
	price float32
}

func NewKeyBoard(p float32) *KeyBoard {
	return &KeyBoard{
		price: p,
	}
}

func (k *KeyBoard) GetPrice() float32 {
	return k.price
}
