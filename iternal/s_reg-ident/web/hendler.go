package web

import (
	//"errors"
	"fmt"
	"log"
	"net/http"
	"pet/iternal/s_reg-ident/db/comsql"
	"pet/iternal/s_reg-ident/str/authe"
	"pet/iternal/s_reg-ident/str/regin"
	"pet/iternal/s_reg-ident/str/salt"
	"pet/pkg/myerr"
	"pet/pkg/pars"
)

func handlerRegPage(w http.ResponseWriter, r *http.Request) {
	log.Println("connection reg")

	err := pars.ParsPage(w, "ui/HTML/reg3.html")
	if err != nil {
		myerr.ServesError(w, err)
		return
	}
}

func (con *Connect) handlerRegProcess(w http.ResponseWriter, r *http.Request) {
	log.Println("process reg")

	if len(r.FormValue("name")) == 0 || len(r.FormValue("password")) == 0 || len(r.FormValue("email")) == 0 {
		err := fmt.Errorf("not se value")
		myerr.ServesError(w, err)
		return
	}

	err := comsql.CheckUinquenessLogin(con.MySQL, r)
	if err != nil {
		myerr.ServesError(w, err)
		return
	}

	err = comsql.CheckUinquenessEmail(con.MySQL, r)
	if err != nil {
		myerr.ServesError(w, err)
		return
	}

	rd, err := regin.New(r)
	if err != nil {
		myerr.ServesError(w, err)
		return
	}

	salt := salt.GenerateSalt()
	key := salt.GeneraterKey(rd.GetPass())
	fmt.Println(*salt, *rd, key)
	err = comsql.SendRegData(con.MySQL, salt, rd, key)
	if err != nil {
		myerr.ServesError(w, err)
		return
	}

	pars.ParsPage(w, "ui/HTML/regstat2.html")
	if err != nil {
		myerr.ServesError(w, err)
		return
	}
}

func handlerAuthPage(w http.ResponseWriter, r *http.Request) {
	log.Println("connection auth")

	err := pars.ParsPage(w, "ui/HTML/auth.html")
	if err != nil {
		myerr.ServesError(w, err)
		return
	}
}

func (con *Connect) handlerAuthProcess(w http.ResponseWriter, r *http.Request) {
	log.Println("process auth")

	if r.Method != "POST" {
		err := fmt.Errorf("method not post")
		myerr.ServesError(w, err)
		return
	}
	if len(r.FormValue("name")) == 0 || len(r.FormValue("password")) == 0 {
		err := fmt.Errorf("not se value")
		myerr.ServesError(w, err)
		return
	}

	resolt, err := comsql.GetAccountData(con.MySQL, r.FormValue("name"))
	if err != nil {
		myerr.ServesError(w, err)
		return
	}

	a := authe.New(resolt, r)
	a.Compare()
	token, err := a.CreateJWT(con.K)
	if err != nil {
		myerr.ServesError(w, err)
		return
	}

	fmt.Print(token)

	pars.ParsPage(w, "ui/HTML/regstat2.html")
	if err != nil {
		log.Println(err)
	}
}
