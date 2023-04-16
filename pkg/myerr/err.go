package myerr

import (
	"html/template"
	"log"
	"net/http"
)

func ServesError(w http.ResponseWriter, err error) {
	log.Println(err)
	w.WriteHeader(http.StatusMethodNotAllowed)
	templ, errs := template.ParseFiles("ui/HTML/regstat2.html")
	if errs != nil {
		log.Println("sistem err: ", errs)
	}
	templ.Execute(w, err)
}
