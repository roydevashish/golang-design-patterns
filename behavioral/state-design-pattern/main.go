package main

import (
	"github.com/roydevashish/golang-design-patterns/state/document"
	userrole "github.com/roydevashish/golang-design-patterns/state/types/user_role"
)

func main() {
	doc := document.NewDocumet(userrole.READER)
	doc.GetState() // Draft

	doc.Publish()
	doc.GetState() // Draft

	doc.SetCurrentUserRole(userrole.EDITOR)
	doc.Publish()
	doc.GetState() // Moderation

	doc.Publish()
	doc.GetState() // Moderation

	doc.SetCurrentUserRole(userrole.ADMIN)
	doc.Publish()
	doc.GetState() // Publish
}
