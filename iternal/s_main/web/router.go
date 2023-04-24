package web

import (
	"crypto"
	"net/http"
	"pet/pkg/pars"
)

type Connect struct {
	PageKesh  *pars.HashTempl
	PublicKey *crypto.PublicKey
	UrlServer *string
}

func (con *Connect) Rout() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/main", con.ValidUrl(con.hendlerMain))
	return mux
}
