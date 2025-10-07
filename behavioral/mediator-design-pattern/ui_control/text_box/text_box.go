package textbox

import dialogbox "github.com/roydevashish/golang-design-patterns/mediator/dialog_box"

type TextBox struct {
	owner dialogbox.DialogBox
	text  string
}

func NewTextBox(owner dialogbox.DialogBox) *TextBox {
	return &TextBox{owner: owner}
}

func (t *TextBox) Get() string {
	return t.text
}

func (t *TextBox) Set(text string) {
	t.text = text
	t.owner.Changed(t)
}
