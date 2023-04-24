package jwtpkg

import (
	"crypto"

	"github.com/golang-jwt/jwt"
)

func VerificationJWTAcc(tokenstring string, key *crypto.PublicKey) (token *jwt.Token, err error) {
	token, err = jwt.Parse(tokenstring, func(token *jwt.Token) (interface{}, error) {
		return *key, nil
	})
	if err != nil {
		return token, err
	}
	return token, nil
}
