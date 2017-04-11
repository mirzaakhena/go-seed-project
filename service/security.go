package service

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var key = []byte{
	'\x4a', '\x21', '\x77', '\xdf', '\xe1', '\x8b', '\xcc', '\x78',
	'\xb1', '\x95', '\x8e', '\x7d', '\x22', '\xdb', '\x08', '\x29',
	'\xa1', '\x28', '\x29', '\xd4', '\xbd', '\x12', '\xe5', '\xda',
	'\xa0', '\xf5', '\x4f', '\x36', '\x47', '\xd4', '\xd4', '\x5c',

	'\x7c', '\xea', '\xbd', '\x0d', '\x6f', '\x2a', '\xec', '\x2d',
	'\x9a', '\xa7', '\x3a', '\x92', '\x81', '\x00', '\xe3', '\xa0',
	'\x93', '\x0b', '\x10', '\x75', '\xc6', '\x97', '\xa8', '\x6a',
	'\xc9', '\xfd', '\x48', '\xdc', '\xbc', '\xb5', '\xce', '\xff',
}

type CustomJwt struct {
	jwt.StandardClaims
	UserId string `json:"userId"`
}

func GenerateToken(userId string) (string, error) {

	t := time.Now().UTC()
	claims := &CustomJwt{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			Subject:   "login",
			Issuer:    "miranc",
			Audience:  "mirzaakhena.com",
			IssuedAt:  t.Unix(),
			ExpiresAt: t.Add(24 * time.Hour).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(key)
	if err != nil {
		log.Error(err.Error())
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(ts string) (*CustomJwt, error) {
	token, err := jwt.ParseWithClaims(ts, &CustomJwt{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})

	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	claims, ok := token.Claims.(*CustomJwt)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	if claims.Subject != "login" {
		return nil, errors.New("invalid subject")
	}

	return claims, nil
}
