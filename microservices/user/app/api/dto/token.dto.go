package dto

import user_domain "user/app/domain/user"

type TokenPayload struct {
	UserId string
	Role   user_domain.UserRole
}
