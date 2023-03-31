package auth

import (
	"pet/iternal/s_reg-ident/str/salt"
)

type Auth struct {
	logname  string
	password string
	salt     *salt.Salt
	key      string
}
