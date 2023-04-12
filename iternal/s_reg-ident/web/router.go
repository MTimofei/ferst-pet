package web

import (
	"database/sql"
	"net/http"
	"pet/iternal/s_reg-ident/jwt/re"
	"pet/iternal/s_reg-ident/web/urlcheck"
)

type Connect struct {
	MySQL *sql.DB
	//PostgraSQL *sql.DB
	K *re.Key
}

func (con *Connect) Router() (mux *http.ServeMux) {
	mux = http.NewServeMux()
	mux.HandleFunc("/", urlcheck.CheckURL(con.handlerMain))
	mux.HandleFunc("/reg", urlcheck.CheckURL(handlerRegPage))
	mux.HandleFunc("/reg/process", urlcheck.CheckURL(con.handlerRegProcess))
	mux.HandleFunc("/auth", urlcheck.CheckURL(handlerAuthPage))
	mux.HandleFunc("/auth/process", urlcheck.CheckURL(con.handlerAuthProcess))
	return mux
}
