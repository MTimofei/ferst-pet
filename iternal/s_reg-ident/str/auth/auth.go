package auth

import (
	"net/http"
	"pet/iternal/s_reg-ident/str/account"
	"pet/pkg/convert"
	"reflect"
)

type Auth struct {
	password string
	authdata *account.Account
}

func New(account *account.Account, r *http.Request) (a *Auth) {
	a = &Auth{
		password: r.FormValue("password"),
		authdata: account,
	}
	return a
}

func (a *Auth) Compare() (b bool) {
	inkey := a.authdata.Saltauth.GeneraterKey([]byte(a.password))
	dbkey := convert.StrToByte(a.authdata.Key)

	if reflect.DeepEqual(inkey, dbkey) {
		return true
	} else {
		return false
	}

}
