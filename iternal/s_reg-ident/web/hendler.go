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
	rd, err := regin.New(r)
	if err != nil {
		help.ServesError(w, err)
		return
	}
	salt := salt.GenerateSalt()
	key := salt.GeneraterKey(rd.GetPass())
	fmt.Printf("%v", key)

	fmt.Fprintf(w, "%s", http.StatusText(http.StatusOK))
	w.WriteHeader(http.StatusOK)
}
