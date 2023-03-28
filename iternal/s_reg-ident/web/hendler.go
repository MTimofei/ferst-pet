package web

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"pet/pkg/myerr"

	"pet/iternal/s_reg-ident/str/regin"
	"pet/iternal/s_reg-ident/str/regout"
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

func handlerPost(w http.ResponseWriter, r *http.Request) {
	rd, err := regin.New(r)
	if err != nil {
		myerr.ServesError(w, err)
		return
	}
	salt := salt.GenerateSalt()
	key := salt.GeneraterKey(rd.GetPass())
	//fmt.Printf("%v", key)
	regout.NewReg(salt, rd, key)

	fmt.Fprintf(w, "%s", http.StatusText(http.StatusOK))
	w.WriteHeader(http.StatusOK)
}
