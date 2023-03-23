package web

import (
	"net/http"
	"ser_identification/iternal/s_reg-ident/web/urlcheck"
)

func Router() (mux *http.ServeMux) {
	mux = http.NewServeMux()
	mux.HandleFunc("/reg", urlcheck.CheckURL(handlerIdent))
	mux.HandleFunc("/reg/process", urlcheck.CheckURL(handlerPost))
	return mux
}
