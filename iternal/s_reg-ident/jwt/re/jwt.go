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

// type ClaimsStoreg struct {
// 	id   int       `json:"id"`
// 	name string    `json:"name"`
// 	iat  time.Time `json:"iat"`
//	exp  time.Time `json:"exp"`
// 	//"roles":a.roles
// }

func GeneratingEncryptionKeys() (k *KeyRef, err error) {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}
	k = &KeyRef{
		privatekey: key,
	}
	return
}

func (k *KeyRef) CreateJWTRefresh(a *account.Account) (tokenString string, err error) {
	*k.Id++

	tokenN := jwt.New(jwt.SigningMethodES256)
	tokenN.Header["kid"] = *k.Id
	tokenN.Header["name"] = "ref"

	tokenN.Claims = jwt.MapClaims{
		"id":   a.Id,
		"name": a.Logname,
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
		//"roles":a.roles
	}

	tokenString, err = tokenN.SignedString(k.privatekey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (k *KeyRef) VerifiedJWTRef(tokenString string) (token *jwt.Token, err error) {
	token, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return k.privatekey.Public(), nil
	})
	if err != nil {
		return token, err
	}
	return token, nil
}
