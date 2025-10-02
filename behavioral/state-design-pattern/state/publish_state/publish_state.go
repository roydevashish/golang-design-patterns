package publishstate

import (
	documentcontext "github.com/roydevashish/golang-design-patterns/state/types/document_context"
)

type PublishState struct {
	document documentcontext.DocumentContext
}

func NewPublishState(doc documentcontext.DocumentContext) *PublishState {
	return &PublishState{
		document: doc,
	}
}

func (p *PublishState) GetState() string {
	return "publish"
}

func (p *PublishState) Publish() {
	// do nothing
}
