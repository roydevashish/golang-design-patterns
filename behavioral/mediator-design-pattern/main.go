package main

import (
	dialogbox "github.com/roydevashish/golang-design-patterns/mediator/dialog_box"
	postdialogbox "github.com/roydevashish/golang-design-patterns/mediator/dialog_box/post_dialog_box"
)

func main() {
	var dialog dialogbox.DialogBox = postdialogbox.NewPostsDialogBox()
	dialog.SimulateUserInteraction()
}
