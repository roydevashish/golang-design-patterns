package button

import dialogbox "github.com/roydevashish/golang-design-patterns/mediator/dialog_box"

type Button struct {
	owner     dialogbox.DialogBox
	isEnabled bool
}

func NewButton(owner dialogbox.DialogBox) *Button {
	return &Button{owner: owner}
}

func (b *Button) Get() bool {
	return b.isEnabled
}

func (b *Button) Set(isEnabled bool) {
	b.isEnabled = isEnabled
	b.owner.Changed(b)
}
