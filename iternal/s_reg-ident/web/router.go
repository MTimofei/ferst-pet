package web

import (
	"context"
	"database/sql"
	"log"
	"net"
	"net/http"
	"pet/iternal/s_reg-ident/jwt/ac"
	"pet/iternal/s_reg-ident/jwt/re"
	"pet/pkg/pars"
)

type Connect struct {
	MySQL     *sql.DB
	KeyRef    *re.KeyRef
	KeyAcc    *ac.KeyAcc
	KeshTempl *pars.KeshTempl
}

func (con *Connect) StartServe(addr *string) {
	log.Println("ser")
	n := &net.ListenConfig{}
	lis, err := n.Listen(context.Background(), "tcp", *addr)
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		err := http.Serve(lis, con.Router())
		log.Fatal(err)
	}()
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
