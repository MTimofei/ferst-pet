package main

import (
	"flag"
	"log"
	"net/http"
	"ser_identification/iternal/s_reg-ident/web"
)

var (
	addr = flag.String("addr", "localhost:8888", "tcp/ip server")
)

func main() {
	flag.Parse()
	log.Println(*addr)
	log.Fatal(http.ListenAndServe(*addr, web.Router()))
}
