package user_domain

func (user *User) IsAdmin() bool {
	return user.Role == ROLE_ADMIN
}
