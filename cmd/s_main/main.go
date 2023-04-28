package main

import (
	"crypto/rsa"
	"flag"
	"net/http"
	realtime "pet/integration_verification/real-time"
	"pet/integration_verification/synchron"
	"pet/iternal/s_main/web"
	"pet/pkg/myerr"
	"pet/pkg/pars"
)

var (
	block       = make(chan int)
	keytransfer = make(chan *rsa.PublicKey)

	urlserves = "http://localhost:8888/main"

	addr      = flag.String("addr", "localhost:8888", "address serer")
	addrGRPC  = flag.String("addrGRPC", "localhost:8000", "adderss gRPC")
	pathDirUi = flag.String("puth-dir-ui", "ui/HTML/", "puth ui directory")
)

func main() {
	flag.Parse()

	synchron.StartSystem()

	keshtempls := pars.New("hi", "regstat")
	keshtempls.LoadHash(*pathDirUi)

	key := &realtime.PublicKey{}
	con := &web.Connect{
		PageKesh:  &keshtempls,
		PublicKey: key,
		UrlServer: &urlserves,
	}
	realtime.StartUpdate(addrGRPC, key, keytransfer, block)
	myerr.LogFatal(http.ListenAndServe(*addr, con.Rout()))
	<-block
}
