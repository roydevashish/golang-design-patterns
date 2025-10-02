package document

import (
	"fmt"

	"github.com/roydevashish/golang-design-patterns/state/state"
	draftstate "github.com/roydevashish/golang-design-patterns/state/state/draft_state"
	userrole "github.com/roydevashish/golang-design-patterns/state/types/user_role"
)

type Document struct {
	state           state.State
	currentUserRole userrole.UserRole
}

func NewDocumet(userRole userrole.UserRole) *Document {
	doc := &Document{}
	doc.SetState(draftstate.NewDraftState(doc))
	doc.SetCurrentUserRole(userRole)
	return doc
}

func (d *Document) Publish() {
	d.state.Publish()
}

func (d *Document) GetState() {
	fmt.Println(d.state.GetState())
}

func (d *Document) SetState(state state.State) {
	d.state = state
}

func (d *Document) GetCurrentUserRole() userrole.UserRole {
	return d.currentUserRole
}

func (d *Document) SetCurrentUserRole(userRole userrole.UserRole) {
	d.currentUserRole = userRole
}
