package authe

import (
	"errors"
	"net/http"
	"pet/iternal/s_reg-ident/jwt/re"
	"pet/iternal/s_reg-ident/str/account"
	"pet/pkg/convert"
	"reflect"
)

type Authe struct {
	password string
	Authdata *account.Account
	isit     bool
}

func New(account *account.Account, r *http.Request) (a *Authe) {
	a = &Authe{
		password: r.FormValue("password"),
		Authdata: account,
		isit:     false,
	}

	return a
}

func (a *Authe) Compare() {
	inkey := a.Authdata.Saltauth.GeneraterKey([]byte(a.password))
	dbkey := convert.StrToByte(a.Authdata.Key)
	if reflect.DeepEqual(inkey, dbkey) {
		a.isit = true
	} else {
		a.isit = false
	}

}

func (a *Authe) CreateJWT(K *re.Key) (t string, err error) {
	if a.isit {
		t, err = K.CreateJWTRefresh(a.Authdata)
		if err != nil {
			return "", err
		}
		return t, nil
	} else {
		err = errors.New("not valid data")
		return "", err
	}
}
