package web

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"pet/pkg/myerr"

	"pet/iternal/s_reg-ident/db/comsql"
	"pet/iternal/s_reg-ident/str/regin"
	"pet/iternal/s_reg-ident/str/salt"
)

func handlerIdent(w http.ResponseWriter, r *http.Request) {
	log.Println("connection")
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

func (con *ConnectDB) handlerPost(w http.ResponseWriter, r *http.Request) {
	if len(r.FormValue("name")) == 0 || len(r.FormValue("password")) == 0 || len(r.FormValue("email")) == 0 {
		err := fmt.Errorf("not se value")
		myerr.ServesError(w, err)
		return
	}

	err := comsql.CheckUnquenessLogin(con.MySQL, r)
	if err != nil {
		myerr.ServesError(w, err)
		return
	}

	err = comsql.CheckUnquenessEmail(con.MySQL, r)
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
	fmt.Fprintf(w, "%s", http.StatusText(http.StatusOK))
	w.WriteHeader(http.StatusOK)
}
