package web

import (
	"net/http"
	"pet/pkg/pars"
)

type Connect struct {
	PageHash *pars.HashTempl
}

func (con *Connect) Rout() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/main", con.ValidUrl(con.hendlerMain))
	return mux
}
