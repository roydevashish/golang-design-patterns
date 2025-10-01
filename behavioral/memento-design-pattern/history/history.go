// Caretaker

package history

import (
	"fmt"
	"slices"

	"github.com/roydevashish/golang-design-patterns/memento/editor"
	editorstate "github.com/roydevashish/golang-design-patterns/memento/editorState"
)

type History struct {
	state  []editorstate.EditorState
	editor *editor.Editor
}

func New(editor *editor.Editor) *History {
	return &History{
		editor: editor,
	}
}

func (h *History) Backup() {
	h.state = append(h.state, *h.editor.CreateState())
}

func (h *History) Undo() {
	if len(h.state) == 0 {
		return
	}

	prevState := h.state[len(h.state)-1]
	h.state = slices.Delete(h.state, len(h.state)-1, len(h.state))
	h.editor.Restore(&prevState)
}

func (h *History) ShowHistory() {
	fmt.Println("History: Here is the list of mementos ⬇️")
	for _, v := range h.state {
		v.GetState()
	}
}
