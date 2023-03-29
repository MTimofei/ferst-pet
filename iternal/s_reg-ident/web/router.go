package web

import (
	"database/sql"
	"net/http"

	"pet/iternal/s_reg-ident/web/urlcheck"
)

type ConnectDB struct {
	MySQL *sql.DB
	//PostgraSQL *sql.DB
}

func (con *ConnectDB) Router() (mux *http.ServeMux) {
	mux = http.NewServeMux()
	mux.HandleFunc("/reg", urlcheck.CheckURL(handlerIdent))
	mux.HandleFunc("/reg/process", urlcheck.CheckURL(con.handlerPost))
	return mux
}
