package main

import (
	"crypto/rsa"
	"flag"
	"log"
	"net/http"
	realtime "pet/integration_verification/real-time"
	"pet/integration_verification/synchron"
	"pet/iternal/s_main/web"
	"pet/pkg/pars"
)

var (
	nameserves = "main"

	keytransfer = make(chan *rsa.PublicKey)

	urlserves = "http://localhost:8888/main"

	addr      = flag.String("addr", "localhost:8888", "address serer")
	addrGRPC  = flag.String("addrGRPC", "localhost:8000", "adderss gRPC")
	pathDirUi = flag.String("puth-dir-ui", "ui/HTML/", "puth ui directory")
)

func main() {
	flag.Parse()

	synchron.StartSystem(&nameserves)

	keshtempls := pars.New("hi", "regstat")
	keshtempls.LoadHash(*pathDirUi)

	key := &realtime.PublicKey{}
	con := &web.Connect{
		PageKesh:  &keshtempls,
		PublicKey: key,
		UrlServer: &urlserves,
	}
	realtime.StartUpdate(addrGRPC, key, keytransfer)
	log.Fatal(http.ListenAndServe(*addr, con.Rout()))
}
