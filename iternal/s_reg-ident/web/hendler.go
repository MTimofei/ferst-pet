package web

import (
	"fmt"
	"log"
	"net/http"
	"pet/iternal/s_reg-ident/db/comsql"
	"pet/iternal/s_reg-ident/str/auth"
	"pet/iternal/s_reg-ident/str/regin"
	"pet/iternal/s_reg-ident/str/salt"
	"pet/pkg/myerr"
	"pet/pkg/pars"
)

func handlerRegPage(w http.ResponseWriter, r *http.Request) {
	log.Println("connection reg")

	err := pars.ParsPage(w, "ui/HTML/reg3.html")

	// templ, err := template.ParseFiles("ui/HTML/reg3.html")
	// if err != nil {
	// 	myerr.ServesError(w, err)
	// 	return
	// }
	// err = templ.Execute(w, nil)

	if err != nil {
		myerr.ServesError(w, err)
		return
	}
}

func (con *ConnectDB) handlerRegProcess(w http.ResponseWriter, r *http.Request) {
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

func (con *ConnectDB) handlerAuthProcess(w http.ResponseWriter, r *http.Request) {
	log.Println("process auth")

	resolt, err := comsql.GetAccountData(con.MySQL, r.FormValue("name"))
	if err != nil {
		myerr.ServesError(w, err)
		return
	}

	a := auth.New(resolt, r)
	b := a.Compare()
	log.Println(b)
	if !b {
		return
	}

	pars.ParsPage(w, "ui/HTML/regstat2.html")
	if err != nil {
		log.Println(err)
	}
}
