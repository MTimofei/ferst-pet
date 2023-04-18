package ac

import (
	"crypto/rand"
	"crypto/rsa"
	"pet/iternal/s_reg-ident/str/account"
	"time"

	"github.com/golang-jwt/jwt"
)

type KeyAcc struct {
	privatekey *rsa.PrivateKey
	Id         *int64
}

func GenerateRSAKey() (k *KeyAcc, err error) {
	key, err := rsa.GenerateMultiPrimeKey(rand.Reader, 4, 2048)
	if err != nil {
		return nil, err
	}
	k = &KeyAcc{
		privatekey: key,
	}
	return k, nil
}

func (k *KeyAcc) CreateJWTAcc(url string, a *account.Account) (tokenString string) {
	*k.Id++
	token := jwt.New(jwt.SigningMethodRS256)
	token.Header["kid"] = *k.Id
	token.Header["name"] = "acc"
	token.Claims = jwt.MapClaims{
		"id":   a.Id,
		"name": a.Logname,
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
		//"roles":a.roles
	}
	return tokenString
}
