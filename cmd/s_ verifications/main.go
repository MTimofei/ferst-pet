package main

//accaunt_ser fZLma7
import (
	"crypto/ecdsa"
	"crypto/rsa"
	"flag"
	"log"
	"pet/iternal/s_reg-ident/grpcser"
	"pet/iternal/s_reg-ident/jwt/ac"
	"pet/iternal/s_reg-ident/jwt/re"
	realtime "pet/iternal/s_reg-ident/real-time"
	"pet/iternal/s_reg-ident/web"
	"pet/pkg/pars"
	"pet/pkg/sql/mysqlcon"
)

var (
	blockmainfn     = make(chan int)
	transportrefkey = make(chan *ecdsa.PrivateKey)
	transportacckey = make(chan *rsa.PrivateKey)

	idjwtref int64 = 0
	idjwtacc int64 = 0

	addr      = flag.String("addr", ":8889", "adderss server")
	addrGRPC  = flag.String("addrGRPC", "localhost:8000", "adderss gRPC")
	pathDirUi = flag.String("puth-dir-ui", "ui/HTML/", "puth ui directory")
	addrMySQL = flag.String("adder-MySQL", "accaunt_ser:fZLma7@/ppa?parseTime=true", "adderss mysql")
)

func main() {
	// fmt.Printf("Ограничение количества горутин: %d\n", runtime.GOMAXPROCS(0))
	// n := 40
	// fmt.Printf("Установка ограничения количества горутин на %d процессоров\n", n)
	// runtime.GOMAXPROCS(n)

	flag.Parse()

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
		MySQL:     dbcon,
		KeyRef:    keyref,
		KeyAcc:    keyacc,
		KeshTempl: &hesh,
	}

	go realtime.RealTimeGenerateEncryptionKeys(transportrefkey)
	go realtime.RealTimeGenerateRSAKey(transportacckey)

	go realtime.UpdateRefPrivateKey(con, transportrefkey)
	go realtime.UpdateAccPrivateKey(con, transportacckey)

	grpcser.StartServerGRPC(*addrGRPC, keyacc)
	con.StartServe(addr)
	log.Println("ALL READY")
	<-blockmainfn
}
