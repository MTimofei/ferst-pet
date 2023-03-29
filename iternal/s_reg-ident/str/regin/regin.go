package regin

import (
	"fmt"
	"net/http"
)

type RegDataIn struct {
	loginname string
	password  string
	email     string
}

func New(r *http.Request) (*RegDataIn, error) {

	if r.Method != "POST" {
		err := fmt.Errorf("method not post")
		return nil, err
	}

	var rd = &RegDataIn{}
	rd.loginname = r.FormValue("name")
	rd.password = r.FormValue("password")
	rd.email = r.FormValue("email")

	return rd, nil
}

func (rd *RegDataIn) GetPass() (pass []byte) {
	pass = []byte(rd.password)
	return pass
}

func (rd *RegDataIn) GetRDIn() (logname, email string) {
	return rd.loginname, rd.email
}
