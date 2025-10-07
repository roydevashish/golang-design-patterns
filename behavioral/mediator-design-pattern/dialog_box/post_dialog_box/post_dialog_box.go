package postdialogbox

import (
	"fmt"

	"github.com/roydevashish/golang-design-patterns/mediator/ui_control/button"
	listbox "github.com/roydevashish/golang-design-patterns/mediator/ui_control/list_box"
	textbox "github.com/roydevashish/golang-design-patterns/mediator/ui_control/text_box"
)

type PostsDialogBox struct {
	postsListBox *listbox.ListBox
	titleTextBox *textbox.TextBox
	saveButton   *button.Button
}

func NewPostsDialogBox() *PostsDialogBox {
	pdb := &PostsDialogBox{}
	pdb.postsListBox = listbox.NewListBox(pdb)
	pdb.titleTextBox = textbox.NewTextBox(pdb)
	pdb.saveButton = button.NewButton(pdb)
	pdb.saveButton.Set(false) // initially disabled
	return pdb
}

func (p *PostsDialogBox) SimulateUserInteraction() {
	p.postsListBox.Set("Post 2")
	fmt.Println("Title text box: " + p.titleTextBox.Get())
	fmt.Println("Button enabled: ", p.saveButton.Get())

	p.titleTextBox.Set("")
	fmt.Println("Title text box: " + p.titleTextBox.Get())
	fmt.Println("Button enabled: ", p.saveButton.Get())
}

func (p *PostsDialogBox) Changed(uiControl interface{}) {
	if lb, ok := uiControl.(*listbox.ListBox); ok && lb == p.postsListBox {
		p.handlePostChanged()
	} else if tb, ok := uiControl.(*textbox.TextBox); ok && tb == p.titleTextBox {
		p.handleTitleChanged()
	}
}

func (p *PostsDialogBox) handlePostChanged() {
	p.titleTextBox.Set(p.postsListBox.Get())
	p.saveButton.Set(true)
}

func (p *PostsDialogBox) handleTitleChanged() {
	isTitleEmpty := p.titleTextBox.Get() == ""
	p.saveButton.Set(!isTitleEmpty)
}
