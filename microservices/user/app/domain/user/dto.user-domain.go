package user_domain

type CreateUserDto struct {
	Login    string   `json:"login"`
	Password string   `json:"password"`
	Role     UserRole `json:"role"`
}
