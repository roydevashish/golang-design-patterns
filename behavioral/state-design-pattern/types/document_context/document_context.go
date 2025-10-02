package documentcontext

import (
	"github.com/roydevashish/golang-design-patterns/state/state"
	userrole "github.com/roydevashish/golang-design-patterns/state/types/user_role"
)

type DocumentContext interface {
	GetState()
	SetState(state state.State)
	GetCurrentUserRole() userrole.UserRole
	SetCurrentUserRole(userRole userrole.UserRole)
}
