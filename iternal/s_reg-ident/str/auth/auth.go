package auth

import (
	"pet/iternal/s_reg-ident/db/comsql"
	"pet/iternal/s_reg-ident/str/salt"
)

type Auth struct {
	id       int
	logname  string
	password string
	salt     *salt.Salt
	key      string
}

func New(account *comsql.Account) (a *Auth) {
	a.id = account.Id
	a.key = account.Key
	return a
}
