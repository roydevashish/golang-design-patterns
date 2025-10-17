package authentication

type Authentication struct{}

func NewAuthentication() *Authentication {
	return &Authentication{}
}

func (a *Authentication) CheckAuth() bool {
	return true
}
