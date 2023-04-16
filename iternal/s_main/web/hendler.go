package web

import (
	"net/http"
	"pet/pkg/myerr"
	"pet/pkg/pars"
)

func (con *Connect) hendlerMain(w http.ResponseWriter, r *http.Request) {
	err := pars.ParsPage(w, "reg", con.PageHash)
	if err != nil {
		myerr.ServesError(w, con.PageHash, err)
	}
}
