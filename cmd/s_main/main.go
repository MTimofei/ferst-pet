package main

import (
	"flag"
	"log"
	"net/http"
	"pet/integration_auth/grpcclient"
	"pet/iternal/s_main/web"
	"pet/pkg/myerr"
	"pet/pkg/pars"
)

var (
	urlserves = "http://localhost:8888/main"

	addr      = flag.String("addr", "localhost:8888", "address serer")
	addrGRPC  = flag.String("addrGRPC", "localhost:8000", "adderss gRPC")
	pathDirUi = flag.String("puth-dir-ui", "ui/HTML/", "puth ui directory")
)

func main() {
	keshtempls := pars.New("hi", "regstat")
	keshtempls.LoadHash(*pathDirUi)

	key, err := grpcclient.ResponsGRPC(addrGRPC)

	if err != nil {
		log.Println(err)
	}

	con := web.Connect{
		PageKesh:  &keshtempls,
		PublicKey: &key,
		UrlServer: &urlserves,
	}
	myerr.LogFatal(http.ListenAndServe(*addr, con.Rout()))
}
