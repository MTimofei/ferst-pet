package auth

import (
	"errors"
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

func (a *Auth) Compare() (b bool, err error) {
	inkey := a.authdata.Saltauth.GeneraterKey([]byte(a.password))
	dbkey, err := convert.StrToByte(a.authdata.Key)
	if err != nil {
		return false, err
	}
	if reflect.DeepEqual(inkey, dbkey) {
		return true, nil
	} else {
		err = errors.New("invalid password")
		return false, err
	}

}
