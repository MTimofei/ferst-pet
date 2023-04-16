package main

import (
	"flag"
	"log"
	"net/http"
	"pet/iternal/s_main/web"
)

var (
	addr = flag.String("addr", "lacalhost:8888", "address serer")
)

func main() {
	err := http.ListenAndServe(*addr, web.Rout())
	if err != nil {
		log.Panicln(err)
	}
}
