package main

import (
	"flag"
	"log"
	"net/http"
	"pet/iternal/s_reg-ident/web"
)

var (
	addr = flag.String("addr", "localhost:8889", "tcp/ip server")
)

func main() {
	flag.Parse()
	log.Fatal(http.ListenAndServe(*addr, web.Router()))
}
