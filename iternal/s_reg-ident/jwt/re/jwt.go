package re

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"pet/iternal/s_reg-ident/str/account"
	"time"

	"github.com/golang-jwt/jwt"
)

type KeyRef struct {
	privatekey *ecdsa.PrivateKey
	Id         *int64
}

func GeneratingEncryptionKeys() (key *KeyRef, err error) {
	privatekey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}
	key = &KeyRef{
		privatekey: privatekey,
	}
	return key, nil
}

func (key *KeyRef) CreateJWTRefresh(account *account.Account) (tokenString string, err error) {
	*key.Id++

	token := jwt.New(jwt.SigningMethodES256)
	token.Header["kid"] = *key.Id
	token.Header["name"] = "ref"

	token.Claims = jwt.MapClaims{
		"id":   account.Id,
		"name": account.Logname,
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
		//"roles":a.roles
	}

	tokenString, err = token.SignedString(key.privatekey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (key *KeyRef) VerifiedJWTRef(tokenString string) (token *jwt.Token, err error) {
	token, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return key.privatekey.Public(), nil
	})
	if err != nil {
		return token, err
	}
	return token, nil
}
