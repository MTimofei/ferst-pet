package pars

import (
	"html/template"
	"net/http"
	"pet/pkg/myerr"
)

func ParsPage(w http.ResponseWriter, path string) {
	templ, err := template.ParseFiles(path)
	if err != nil {
		myerr.ServesError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	err = templ.Execute(w, http.StatusText(http.StatusOK))
	if err != nil {
		myerr.ServesError(w, err)
		return
	}

}
