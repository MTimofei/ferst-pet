package web

import (
	"html/template"
	"log"
	"net/http"
	"ser_identification/iternal/s_reg-ident/key"
	"ser_identification/iternal/s_reg-ident/str/reg"
)

func handlerIdent(w http.ResponseWriter, r *http.Request) {
	log.Println("connection")
	templ, err := template.ParseFiles("ui/HTML/reg3.html")
	if err != nil {
		log.Println(err)
		return
	}
	err = templ.Execute(w, nil)
	if err != nil {
		log.Println(err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func handlerPost(w http.ResponseWriter, r *http.Request) {
	log.Println("data transmission")
	dynamic := key.GenerateDenamicSolt()
	rd, err := reg.New(r)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		log.Println(err)
		return
	}
	rd.PreparationRegistratDate(dynamic)

	w.WriteHeader(http.StatusOK)
}
