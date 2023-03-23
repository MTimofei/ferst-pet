package reg

import (
	"fmt"
	"net/http"
	"ser_identification/iternal/s_reg-ident/key"
)

// ////////////////////////
type regData struct {
	logname     string
	password    []byte
	dynamicsalt []byte
	email       string
}

//func New(log, pas, email string) (rd regData) {
//	rd.logname = log
//	rd.password = pas
//	rd.email = email
//	return rd
//}

func New(r *http.Request) (*regData, error) {
	if len(r.FormValue("name")) == 0 || len(r.FormValue("password")) == 0 || len(r.FormValue("email")) == 0 {
		err := fmt.Errorf("not se value")
		return nil, err
	}
	var rd = &regData{}
	rd.logname = r.FormValue("name")
	rd.password = []byte(r.FormValue("password"))
	rd.email = r.FormValue("email")

	return rd, nil
}

func (rd *regData) PreparationRegistratDate(dynamicsalt []byte) {
	pashash := key.GeneraterKey(dynamicsalt, rd.password)
	rd.password = pashash
	rd.dynamicsalt = dynamicsalt
}
