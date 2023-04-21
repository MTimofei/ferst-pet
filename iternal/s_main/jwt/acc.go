package jwt

import (
	"crypto"

	"github.com/golang-jwt/jwt"
)

func VerificationJWTAcc(tokenstring string, k *crypto.PublicKey) (token *jwt.Token, err error) {
	token, err = jwt.Parse(tokenstring, func(token *jwt.Token) (interface{}, error) {
		return k, nil
	})
	if err != nil {
		return token, err
	}
	return token, nil
}
