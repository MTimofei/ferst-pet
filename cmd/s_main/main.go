package main

import (
	"flag"
	"log"
	"net/http"
	"pet/iternal/s_main/web"
	"pet/pkg/pars"
)

// crypto/ecdsa: verification error
var (
	addr      = flag.String("addr", "localhost:8888", "address serer")
	pathDirUi = flag.String("puth-dir-ui", "ui/HTML/", "puth ui directory")
)

func main() {
	h := pars.New("hi", "regstat")
	h.LoadHash(*pathDirUi)
	con := web.Connect{
		PageHash: &h,
	}
	err := http.ListenAndServe(*addr, con.Rout())
	if err != nil {
		log.Panicln(err)
	}
}
