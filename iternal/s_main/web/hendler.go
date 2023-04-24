package web

import (
	"net/http"
	"pet/integration_auth/cookie"
	"pet/iternal/s_main/str"
	"pet/pkg/myerr"
	"pet/pkg/pars"
)

func (con *Connect) hendlerMain(w http.ResponseWriter, r *http.Request) {
	token, err := cookie.HandlingCookiesClients(con.UrlServer, con.PageKesh, con.PublicKey, w, r)

	err = pars.ParsPage(w, "hi", con.PageKesh, str.NewAccountFromJWT(token))
	if err != nil {
		myerr.ServesError(w, con.PageKesh, err)
		return
	}
}
