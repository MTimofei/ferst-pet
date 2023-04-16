package main

import (
	"flag"
	"log"
	"net/http"
	"pet/iternal/s_main/web"
	"pet/pkg/pars"
)

var (
	addr = flag.String("addr", "localhost:8888", "address serer")
)

func main() {
	h := pars.New("reg", "auth", "regstat")
	h.LoadHash("ui/HTML/")
	con := &web.Connect{
		PageHash: &h,
	}
	err := http.ListenAndServe(*addr, con.Rout())
	if err != nil {
		log.Panicln(err)
	}
}
