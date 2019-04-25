package util

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"time"
)

var (
	ErrSigningKeyCannotEmpty = errors.New("signing key can't be empty")
	ErrTokenNotValid = errors.New("token not valid")
)

type UserClaims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

var mySigningKey = []byte("hello hole hole")

func InitSigningKey(key []byte) error {
	if len(key) == 0 {
		return ErrSigningKeyCannotEmpty
	}
	mySigningKey = key

	return nil
}

// GenUserJwtToken 生成用户认证 token
func GenUserJwtToken(uID string, expireDur time.Duration) (string, error) {
	// Create the Claims
	claims := &UserClaims{
		UserID:   uID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expireDur).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}
	return ss, nil
}

// ValidUserJwtToken 验证是否为正确的 token
func ValidUserJwtToken(token string) (uID string, err error) {
	var claims UserClaims
	jt, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (i interface{}, e error) {
		return mySigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if !jt.Valid {
		return "", ErrTokenNotValid
	}

	return claims.UserID, nil
}