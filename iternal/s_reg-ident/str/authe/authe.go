package authe

import (
	"errors"
	"net/http"
	"pet/iternal/s_reg-ident/jwt/re"
	"pet/iternal/s_reg-ident/str/account"
	"pet/pkg/convert"
	"reflect"

	"github.com/golang-jwt/jwt"
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

func nweAuthTrueFromJWT(token *jwt.Token) (a *Authe) {
	claims := token.Claims.(jwt.MapClaims)
	id := claims["id"].(float64)
	name := claims["name"].(string)

	account := account.New(int(id), name, "", "")
	a = &Authe{
		password: "",
		Authdata: account,
		isit:     true,
	}
	return a
}

func AuthRefJWT(k *re.KeyRef, tokenString string) (a *Authe, err error) {
	var token *jwt.Token
	if token, err = k.VerifiedJWTRef(tokenString); err != nil {
		return nil, err
	}
	a = nweAuthTrueFromJWT(token)
	return a, nil

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

func (a *Authe) CreateJWTRef(K *re.KeyRef) (token string, err error) {
	if a.isit {
		token, err = K.CreateJWTRefresh(a.Authdata)
		if err != nil {
			return "", err
		}
		return token, nil
	} else {
		err = errors.New("not valid data")
		return "", err
	}
}
