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
	help.ServesError(w, err)
	err = templ.Execute(w, nil)
	help.ServesError(w, err)
	w.WriteHeader(http.StatusOK)
}

func handlerPost(w http.ResponseWriter, r *http.Request) {
	log.Println("data transmission")
	//salt := key.GenerateDenamicSolt()
	rd, err := regin.New(r)
	help.ServesError(w, err)
	salt := salt.GenerateSalt()
	key := salt.GeneraterKey(rd.GetPass())
	help.ServesError(w, err)
	fmt.Printf("%b", key)
	//if err != nil {
	//	w.WriteHeader(http.StatusServiceUnavailable)
	//	log.Println(err)
	//	return
	//}
	//rd.PreparationRegistratDate(salt)

	w.WriteHeader(http.StatusOK)
}
