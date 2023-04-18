package web

import (
	"database/sql"
	"net/http"
	"pet/iternal/s_reg-ident/jwt/ac"
	"pet/iternal/s_reg-ident/jwt/re"
	"pet/pkg/pars"
)

type Connect struct {
	MySQL     *sql.DB
	KRef      *re.KeyRef
	KAcc      *ac.KeyAcc
	HashTempl *pars.HashTempl
}

func (con *Connect) Router() (mux *http.ServeMux) {
	mux = http.NewServeMux()
	mux.HandleFunc("/", con.ValidUrl(con.handlerMain))
	mux.HandleFunc("/reg", con.ValidUrl(con.handlerRegPage))
	mux.HandleFunc("/auth", con.ValidUrl(con.handlerAuthPage))
	mux.HandleFunc("/reg/process", con.ValidUrl(con.handlerRegProcess))
	mux.HandleFunc("/auth/process", con.ValidUrl(con.handlerAuthProcess))
	return mux
}
