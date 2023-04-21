package str

import (
	"github.com/golang-jwt/jwt"
)

type Account struct {
	Id   int64
	Name string
}

func NewAccountFromJWT(token *jwt.Token) *Account {
	claims := token.Claims.(jwt.MapClaims)
	id := claims["id"].(int64)
	name := claims["name"].(string)

	var a = &Account{
		Id:   id,
		Name: name,
	}
	return a
}
