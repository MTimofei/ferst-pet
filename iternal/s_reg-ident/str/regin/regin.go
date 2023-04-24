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

func New(r *http.Request) (regdatanewuser *RegDataIn, err error) {

	if r.Method != "POST" {
		err := fmt.Errorf("method not post")
		return nil, err
	}

	regdatanewuser = &RegDataIn{}
	regdatanewuser.loginname = r.FormValue("name")
	regdatanewuser.password = r.FormValue("password")
	regdatanewuser.email = r.FormValue("email")

	return regdatanewuser, nil
}

func (regdatanewuser *RegDataIn) GetPass() (passwordbytes []byte) {
	passwordbytes = []byte(regdatanewuser.password)
	return passwordbytes
}

func (regdatanewuser *RegDataIn) GetRDIn() (logname, email string) {
	return regdatanewuser.loginname, regdatanewuser.email
}
