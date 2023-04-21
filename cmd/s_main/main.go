package main

import (
	"flag"
	"log"
	"net/http"
	"pet/iternal/s_main/grpcclient"
	"pet/iternal/s_main/web"
	"pet/pkg/myerr"
)

// crypto/ecdsa: verification error
var (
	addr      = flag.String("addr", "localhost:8888", "address serer")
	addrGRPC  = flag.String("addrGRPC", "localhost:8000", "adderss gRPC")
	pathDirUi = flag.String("puth-dir-ui", "ui/HTML/", "puth ui directory")
)

func main() {
	//	h := pars.New("hi", "regstat")
	//h.LoadHash(*pathDirUi)

	k, err := grpcclient.ConnectionGRPC(*addrGRPC)

	if err != nil {
		log.Println(err)
	}

	con := web.Connect{
		//PageHash:  &h,
		PublicKey: &k,
	}
	myerr.LogFatal(http.ListenAndServe(*addr, con.Rout()))
}
