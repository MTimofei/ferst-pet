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

type payloadc struct {
	id   int
	name string
	//roles string
	iat   time.Time
	exp   time.Time
	nonce uint8
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
	ran := []byte{}
	rand2.Seed(time.Now().UnixNano())
	randNum := uint8(rand2.Intn(255))
	ran = append(ran, randNum)
	tokenN := jwt.New(jwt.SigningMethodES256)
	tokenN.Header["kid"] = idkey
	tokenN.Header["name"] = "ref"

	tokenN.Claims = jwt.MapClaims{
		"jti":  a.Id,
		"name": a.Logname,
		//"roles":a.roles
		"iat":   time.Now().Unix(),
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
		"nonce": base64.StdEncoding.EncodeToString(ran),
	}

	tokenString, err = tokenN.SignedString(k.privatekey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
