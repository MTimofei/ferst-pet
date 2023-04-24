package myerr

import (
	"log"
	"net/http"
	"pet/pkg/pars"
)

func ServesError(w http.ResponseWriter, hesh *pars.KeshTempl, err error) {
	log.Println(err)
	w.WriteHeader(http.StatusMethodNotAllowed)

	pars.ParsPage(w, "regstat", hesh, err)

}

func LogFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
