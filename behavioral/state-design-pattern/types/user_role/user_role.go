package userrole

type UserRole string

const (
	READER UserRole = "reader"
	EDITOR UserRole = "editor"
	ADMIN  UserRole = "admin"
)
