package main

//accaunt_ser fZLma7
import (
	"crypto/ecdsa"
	"crypto/rsa"
	"flag"
	"log"
	"pet/integration_verification/synchron"
	"pet/iternal/s_reg-ident/grpcser"
	"pet/iternal/s_reg-ident/jwt/ac"
	"pet/iternal/s_reg-ident/jwt/re"
	realtime "pet/iternal/s_reg-ident/real-time"
	"pet/iternal/s_reg-ident/web"
	"pet/pkg/pars"
	"pet/pkg/sql/mysqlcon"
	"runtime"
)

var (
	block           = make(chan int)
	transportrefkey = make(chan *ecdsa.PrivateKey)
	transportacckey = make(chan *rsa.PrivateKey)

	idjwtref   int64  = 0
	idjwtacc   int64  = 0
	nameserves string = "verifications"

	addr      = flag.String("addr", ":8889", "adderss server")
	addrGRPC  = flag.String("addrGRPC", "localhost:8000", "adderss gRPC")
	pathDirUi = flag.String("puth-dir-ui", "ui/HTML/", "puth ui directory")
	addrMySQL = flag.String("adder-MySQL", "accaunt_ser:fZLma7@/ppa?parseTime=true", "adderss mysql")
)

func main() {
	runtime.GOMAXPROCS(8)
	flag.Parse()

	synchron.StartSystem(&nameserves)

	dbcon, err := mysqlcon.OpenMySQLDB(addrMySQL)
	if err != nil {
		log.Fatal(err)
	}
	defer dbcon.Close()

	keyref := &re.KeyRef{
		Id: &idjwtref,
	}
	keyacc := &ac.KeyAcc{
		Id: &idjwtacc,
	}

	hesh := pars.New("reg", "auth", "regstat", "hi")
	hesh.LoadHash(*pathDirUi)

	con := &web.Connect{
		MySQL:  dbcon,
		KeyRef: keyref,
		KeyAcc: keyacc,
		// KeshTempl: &hesh,
	}

	realtime.StartUpdataKey(con, transportrefkey, transportacckey)
	grpcser.StartServerGRPC(*addrGRPC, keyacc)
	con.StartServe(addr)
	log.Println("ALL READY")
	<-block
}
