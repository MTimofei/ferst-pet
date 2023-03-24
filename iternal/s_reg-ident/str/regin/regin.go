package regin

import (
	"fmt"
	"net/http"
)

type regDataIn struct {
	loginname string
	password  string
	email     string
}

func New(r *http.Request) (*regDataIn, error) {
	if len(r.FormValue("name")) == 0 || len(r.FormValue("password")) == 0 || len(r.FormValue("email")) == 0 {
		err := fmt.Errorf("not se value")
		return nil, err
	}
	var rd = &regDataIn{}
	rd.loginname = r.FormValue("name")
	rd.password = r.FormValue("password")
	rd.email = r.FormValue("email")

	return rd, nil
}

//
//func (rd *regDataIn) PreparationRegistratDate(dynamicsalt []byte) {
//	pashash := key.GeneraterKey(dynamicsalt, rd.password)
//	rd.password = pashash
//	rd.dynamicsalt = dynamicsalt
//}
