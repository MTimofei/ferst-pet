package ac

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"pet/iternal/s_reg-ident/str/account"
	"sync"
	"time"

	"github.com/golang-jwt/jwt"
)

type KeyAcc struct {
	privatekey *rsa.PrivateKey
	Id         *int64
}

func GenerateRSAKey() (privatekey *rsa.PrivateKey, err error) {
	privatekey, err = rsa.GenerateMultiPrimeKey(rand.Reader, 4, 2048)
	if err != nil {
		return nil, err
	}
	return privatekey, nil
}

func (key *KeyAcc) CreateJWTAcc(a *account.Account) (tokenString string, err error) {
	*key.Id++
	token := jwt.New(jwt.SigningMethodRS256)
	token.Header["kid"] = *key.Id
	token.Header["name"] = "acc"
	token.Claims = jwt.MapClaims{
		"id":   a.Id,
		"name": a.Logname,
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

func (key *KeyAcc) GetPublicKey() []byte {
	keybyts := x509.MarshalPKCS1PublicKey(&key.privatekey.PublicKey)
	return keybyts
}

func (key *KeyAcc) Update(privateacckey *rsa.PrivateKey) {
	mutex := sync.Mutex{}
	mutex.Lock()
	key.privatekey = privateacckey
	mutex.Unlock()
}
