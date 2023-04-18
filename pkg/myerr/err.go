package myerr

import (
	"log"
	"net/http"
	"pet/pkg/pars"
)

func ServesError(w http.ResponseWriter, hesh *pars.HashTempl, err error) {
	log.Println(err)
	w.WriteHeader(http.StatusMethodNotAllowed)

	pars.ParsPage(w, "regstat", hesh, err)

}
