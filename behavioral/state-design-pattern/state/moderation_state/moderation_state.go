package moderationstate

import (
	publishstate "github.com/roydevashish/golang-design-patterns/state/state/publish_state"
	documentcontext "github.com/roydevashish/golang-design-patterns/state/types/document_context"
	userrole "github.com/roydevashish/golang-design-patterns/state/types/user_role"
)

type ModerationState struct {
	document documentcontext.DocumentContext
}

func NewModerationState(doc documentcontext.DocumentContext) *ModerationState {
	return &ModerationState{
		document: doc,
	}
}

func (m *ModerationState) GetState() string {
	return "moderation"
}

func (m *ModerationState) Publish() {
	if m.document.GetCurrentUserRole() == userrole.ADMIN {
		m.document.SetState(publishstate.NewPublishState(m.document))
	}
}
