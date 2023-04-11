package re

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/base64"
	rand2 "math/rand"
	"pet/iternal/s_reg-ident/str/account"
	"time"

	"github.com/golang-jwt/jwt"
	//"github.com/golang-jwt/jwt/v5"
)

var idkey int64 = 0

type Key struct {
	privatekey *ecdsa.PrivateKey
}

func GeneratingEncryptionKeys() (k *Key, err error) {
	p, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}
	k = &Key{
		privatekey: p,
	}
	return
}

func (k *Key) CreateJWTRefresh(a *account.Account) (tokenString string, err error) {
	idkey++
	r := []byte{}
	rand2.Seed(time.Now().UnixNano())
	randNum := uint8(rand2.Intn(255))
	r = append(r, randNum)
	token := jwt.New(jwt.SigningMethodES256)
	token.Header["kid"] = idkey

	token.Claims = jwt.MapClaims{
		"id":   a.Id,
		"name": a.Logname,
		//"rights":a.rights
		"iat":   time.Now().Unix(),
		"exp":   time.Now().Add(time.Hour / 60).Unix(),
		"nonce": base64.StdEncoding.EncodeToString(r),
	}

	tokenString, err = token.SignedString(k.privatekey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
