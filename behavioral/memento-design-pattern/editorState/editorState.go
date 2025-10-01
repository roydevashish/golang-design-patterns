// Memento

package editorstate

import (
	"fmt"
	"time"
)

type EditorState struct {
	title          string
	content        string
	stateCreatedAt time.Time
}

func New(title string, content string) *EditorState {
	return &EditorState{
		title:          title,
		content:        content,
		stateCreatedAt: time.Now(),
	}
}

func (e *EditorState) GetTitle() string {
	return e.title
}

func (e *EditorState) GetContent() string {
	return e.content
}

func (e *EditorState) GetDate() time.Time {
	return e.stateCreatedAt
}

func (e *EditorState) GetState() {
	fmt.Printf("%v / title: %q, content: %q \n", e.GetDate(), e.GetTitle(), e.GetContent())
}
