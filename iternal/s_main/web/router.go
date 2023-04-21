package web

import (
	"crypto"
	"net/http"
	"pet/pkg/pars"
)

type Connect struct {
	PageHash  *pars.HashTempl
	PublicKey *crypto.PublicKey
}

func (con *Connect) Rout() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/main", con.ValidUrl(con.hendlerMain))
	return mux
}
