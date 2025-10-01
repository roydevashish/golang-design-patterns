package main

import (
	"fmt"

	"github.com/roydevashish/golang-design-patterns/memento/editor"
	"github.com/roydevashish/golang-design-patterns/memento/history"
)

func main() {
	editor := editor.New()
	history := history.New(editor)

	history.Backup()

	editor.Title = "Test"
	history.Backup()

	editor.Content = "Hello there, my name is Roy"
	history.Backup()

	editor.Title = "The life of a dev: my memento"
	history.Backup()

	fmt.Printf("Title: %v, Content: %v\n", editor.Title, editor.Content)
	history.ShowHistory()

	history.Undo()
	fmt.Printf("Title: %v, Content: %v\n", editor.Title, editor.Content)

	history.Undo()
	fmt.Printf("Title: %v, Content: %v\n", editor.Title, editor.Content)

	history.Undo()
	fmt.Printf("Title: %v, Content: %v\n", editor.Title, editor.Content)

	history.Undo()
	fmt.Printf("Title: %v, Content: %v\n", editor.Title, editor.Content)
}
