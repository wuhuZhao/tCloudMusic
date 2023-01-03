package common

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Claims: jwt struct
type Claims struct {
	username string
	jwt.StandardClaims
}

// TokenExpireDuration: expired time
const TokenExpireDuration = time.Hour * 2

// secrect: token secrect
var secrect = []byte("biztokensecrect")

// GenToken: username to token
func GenToken(username string) (string, error) {
	c := Claims{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "tCloudMusic",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(secrect)
}

// ParseToken: token to username
func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return secrect, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
