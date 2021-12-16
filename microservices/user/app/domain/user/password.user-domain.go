package user_domain

import "golang.org/x/crypto/bcrypt"

func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 6)

	if err != nil {
		return err
	}

	user.Password = string(bytes)

	return nil
}

func (user *User) IsPasswordCorrect(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
