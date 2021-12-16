package dto

type UserRole byte

const (
	ROLE_COMMON UserRole = iota + 1
	ROLE_ADMIN
)
