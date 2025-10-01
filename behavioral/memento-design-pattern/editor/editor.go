// Originator
package editor

import editorstate "github.com/roydevashish/golang-design-patterns/memento/editorState"

type Editor struct {
	Title   string
	Content string
}

func New() *Editor {
	return &Editor{
		Title:   "",
		Content: "",
	}
}

func (e *Editor) CreateState() *editorstate.EditorState {
	return editorstate.New(e.Title, e.Content)
}

func (e *Editor) Restore(state *editorstate.EditorState) {
	e.Title = state.GetTitle()
	e.Content = state.GetContent()
}
