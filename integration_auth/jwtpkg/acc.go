package jwtpkg

import (
	"crypto/rsa"
	"log"

	"github.com/golang-jwt/jwt"
)

func VerificationJWTAcc(tokenstring string, key *rsa.PublicKey) (token *jwt.Token, err error) {
	log.Println(key)
	token, err = jwt.Parse(tokenstring, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return token, err
	}
	return token, nil
}
