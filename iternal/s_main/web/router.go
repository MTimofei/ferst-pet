package web

import (
	"net/http"
	realtime "pet/integration_auth/real-time"
	"pet/pkg/pars"
)

type Connect struct {
	PageKesh  *pars.KeshTempl
	PublicKey *realtime.PublicKey
	UrlServer *string
}

func (con *Connect) Rout() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/main", con.ValidUrl(con.hendlerMain))
	return mux
}
