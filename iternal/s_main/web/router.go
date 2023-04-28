package web

import (
	"log"
	"net/http"
	realtime "pet/integration_verification/real-time"
	"pet/pkg/pars"
)

type Connect struct {
	PageKesh  *pars.KeshTempl
	PublicKey *realtime.PublicKey
	UrlServer *string
}

func (con *Connect) Rout() *http.ServeMux {
	log.Println("ser")
	mux := http.NewServeMux()
	mux.HandleFunc("/main", con.ValidUrl(con.hendlerMain))
	return mux
}
