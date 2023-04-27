package main

import (
	"crypto/rsa"
	"flag"
	"net/http"
	"pet/integration_auth/grpcclient"
	realtime "pet/integration_auth/real-time"
	"pet/iternal/s_main/web"
	"pet/pkg/myerr"
	"pet/pkg/pars"
)

var (
	//blikmainfunc = make(chan int)
	keytransfer = make(chan *rsa.PublicKey)

	urlserves = "http://localhost:8888/main"

	addr      = flag.String("addr", "localhost:8888", "address serer")
	addrGRPC  = flag.String("addrGRPC", "localhost:8000", "adderss gRPC")
	pathDirUi = flag.String("puth-dir-ui", "ui/HTML/", "puth ui directory")
)

func main() {
	keshtempls := pars.New("hi", "regstat")
	keshtempls.LoadHash(*pathDirUi)

	key := &realtime.PublicKey{}
	con := &web.Connect{
		PageKesh:  &keshtempls,
		PublicKey: key,
		UrlServer: &urlserves,
	}
	go grpcclient.RealTimeGetKyeМiaGRPC(addrGRPC, keytransfer)
	go realtime.UpdatePublicKey(key, keytransfer)
	myerr.LogFatal(http.ListenAndServe(*addr, con.Rout()))
}
