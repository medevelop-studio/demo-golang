package user_domain

type UserRole byte

const (
	ROLE_COMMON UserRole = iota + 1
	ROLE_ADMIN
)

type User struct {
	Id    string
	Login string
	Role  UserRole
}
