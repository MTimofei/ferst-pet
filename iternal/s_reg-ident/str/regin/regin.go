package regin

import (
	"fmt"
	"log"
	"net/http"
)

type RegDataIn struct {
	loginname string
	password  string
	email     string
}

func New(r *http.Request) (*RegDataIn, error) {
	log.Println("01data transmission")
	if r.Method != "POST" {
		err := fmt.Errorf("method not post")
		return nil, err
	}
	log.Println("02data transmission")
	if len(r.FormValue("name")) == 0 || len(r.FormValue("password")) == 0 || len(r.FormValue("email")) == 0 {
		err := fmt.Errorf("not se value")
		return nil, err
	}
	log.Println("03data transmission")
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
