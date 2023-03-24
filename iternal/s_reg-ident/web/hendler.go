package web

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"ser_identification/iternal/s_reg-ident/str/regin"
	"ser_identification/iternal/s_reg-ident/str/salt"
	"ser_identification/pkg/help"
)

func handlerIdent(w http.ResponseWriter, r *http.Request) {
	log.Println("connection")
	templ, err := template.ParseFiles("ui/HTML/reg3.html")
	if err != nil {
		help.ServesError(w, err)
		return
	}
	err = templ.Execute(w, nil)
	if err != nil {
		help.ServesError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func handlerPost(w http.ResponseWriter, r *http.Request) {
	log.Println("data transmission")
	rd, err := regin.New(r)
	log.Println("1data transmission")
	if err != nil {
		help.ServesError(w, err)
		return
	}
	log.Println("2data transmission")
	salt := salt.GenerateSalt()
	log.Println("3data transmission")
	key := salt.GeneraterKey(rd.GetPass())
	log.Println("4data transmission")
	fmt.Printf("%v", key)

	w.WriteHeader(http.StatusOK)
}
