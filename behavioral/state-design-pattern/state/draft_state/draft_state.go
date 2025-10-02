package draftstate

import (
	moderationstate "github.com/roydevashish/golang-design-patterns/state/state/moderation_state"
	documentcontext "github.com/roydevashish/golang-design-patterns/state/types/document_context"
	userrole "github.com/roydevashish/golang-design-patterns/state/types/user_role"
)

type DraftState struct {
	document documentcontext.DocumentContext
}

func NewDraftState(doc documentcontext.DocumentContext) *DraftState {
	return &DraftState{
		document: doc,
	}
}

func (d *DraftState) GetState() string {
	return "draft"
}

func (d *DraftState) Publish() {
	if d.document.GetCurrentUserRole() == userrole.EDITOR || d.document.GetCurrentUserRole() == userrole.ADMIN {
		d.document.SetState(moderationstate.NewModerationState(d.document))
	}
}
