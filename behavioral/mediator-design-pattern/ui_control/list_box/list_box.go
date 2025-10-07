package listbox

import dialogbox "github.com/roydevashish/golang-design-patterns/mediator/dialog_box"

type ListBox struct {
	owner     dialogbox.DialogBox
	selection string
}

func NewListBox(owner dialogbox.DialogBox) *ListBox {
	return &ListBox{owner: owner}
}

func (l *ListBox) Get() string {
	return l.selection
}

func (l *ListBox) Set(selection string) {
	l.selection = selection
	l.owner.Changed(l)
}
