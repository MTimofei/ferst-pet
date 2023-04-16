package myerr

import (
	"log"
	"net/http"
	"pet/pkg/pars"
)

func ServesError(w http.ResponseWriter, h *pars.HashTempl, err error) {
	log.Println(err)
	w.WriteHeader(http.StatusMethodNotAllowed)
	templ := h.Hash["regstat"]
	templ.Execute(w, err)
}
