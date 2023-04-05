package web

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"pet/iternal/s_reg-ident/db/comsql"
	"pet/iternal/s_reg-ident/str/regin"
	"pet/iternal/s_reg-ident/str/salt"
	"pet/pkg/myerr"
)

func handlerRegPage(w http.ResponseWriter, r *http.Request) {
	log.Println("connection reg")
	templ, err := template.ParseFiles("ui/HTML/reg3.html")
	if err != nil {
		myerr.ServesError(w, err)
		return
	}
	err = templ.Execute(w, nil)
	if err != nil {
		myerr.ServesError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
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
	//fmt.Printf("%v", key)
	//regout.NewReg(salt, rd, key)

	err = comsql.SendRegData(con.MySQL, salt, rd, key)
	if err != nil {
		myerr.ServesError(w, err)
		return
	}
	templ, err := template.ParseFiles("ui/HTML/regstat2.html")
	if err != nil {
		myerr.ServesError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	err = templ.Execute(w, http.StatusText(http.StatusOK))
	if err != nil {
		log.Panicln(err)
	}
}

func handlerAuthPage(w http.ResponseWriter, r *http.Request) {
	log.Println("connection auth")
	templ, err := template.ParseFiles("ui/HTML/auth.html")
	if err != nil {
		myerr.ServesError(w, err)
		return
	}
	err = templ.Execute(w, nil)
	if err != nil {
		myerr.ServesError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (con *ConnectDB) handlerAuthProcess(w http.ResponseWriter, r *http.Request) {

}
