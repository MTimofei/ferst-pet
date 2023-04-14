package myerr

import (
	"html/template"
	"log"
	"net/http"
)

func ServesError(w http.ResponseWriter, err error) {
	log.Println(err)
	w.WriteHeader(http.StatusMethodNotAllowed)
	templ, err_ := template.ParseFiles("ui/HTML/regstat2.html")
	if err_ != nil {
		log.Println(err_)
	}
	templ.Execute(w, err)
}
