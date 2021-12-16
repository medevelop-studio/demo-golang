package auth_service

import (
	"strings"
	"time"
	user_domain "user/app/domain/user"

	"github.com/dgrijalva/jwt-go"
)

func (service *AuthService) IsTokenValid(tokenString string) bool {
	tokenString = strings.TrimSpace(tokenString)

	if len(tokenString) == 0 {
		return false
	}

	token, err := service.parseToken(tokenString)
	if err != nil {
		return false
	}

	return token.Valid
}

func (service *AuthService) createJWTToken(user *user_domain.User) (string, error) {
	atClaims := jwt.MapClaims{
		"id":   user.Id,
		"role": user.Role,
	}

	atClaims["exp"] = time.Now().Add(time.Millisecond * time.Duration(service.config.TokenLifeTime)).Unix()
	unsignedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	token, err := unsignedToken.SignedString([]byte(service.config.TokenSecret))
	if err != nil {
		return "", err
	}

	return token, nil
}

func (service *AuthService) parseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrUnexpectedMethod
		}
		return []byte(service.config.TokenSecret), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}
