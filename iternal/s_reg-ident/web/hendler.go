package web

import (
	"html/template"
	"log"
	"net/http"
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
	//salt := key.GenerateDenamicSolt()
	//rd, err := regin.New(r)
	//if err != nil {
	//	w.WriteHeader(http.StatusServiceUnavailable)
	//	log.Println(err)
	//	return
	//}
	//rd.PreparationRegistratDate(salt)

	w.WriteHeader(http.StatusOK)
}
