package re

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"log"
	"pet/iternal/s_reg-ident/str/account"
	"time"

	"github.com/golang-jwt/jwt"
	//"github.com/golang-jwt/jwt/v5"
)

type Key struct {
	privatekey *ecdsa.PrivateKey
	Id         *int
}

// type payloadc struct {
// 	id   int
// 	name string
// 	//roles string
// 	iat   time.Time
// 	exp   time.Time
// 	nonce uint8
// }

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
	*k.Id++

	tokenN := jwt.New(jwt.SigningMethodES256)
	tokenN.Header["kid"] = *k.Id
	tokenN.Header["name"] = "ref"

	tokenN.Claims = jwt.MapClaims{
		"jti":  a.Id + *k.Id,
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

func (k *Key) VerifiedJWTRef(tokenString string) error {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return k.privatekey.Public(), nil
	})
	if err != nil {
		return err
	}
	log.Println(token.Valid)

	// log.Println(token.Header)
	// log.Println(token.Claims)
	// log.Println(token.Signature)

	return nil
}
