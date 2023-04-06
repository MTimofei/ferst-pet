package pars

import (
	"html/template"
	"net/http"
)

func ParsPage(w http.ResponseWriter, path string) (err error) {
	templ, err := template.ParseFiles(path)
	if err != nil {
		return err
	}
	w.WriteHeader(http.StatusOK)
	err = templ.Execute(w, http.StatusText(http.StatusOK))
	if err != nil {
		return err
	}
	return nil
}
