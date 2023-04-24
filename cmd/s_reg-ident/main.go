package main

//accaunt_ser fZLma7
import (
	"flag"
	"log"
	"pet/iternal/s_reg-ident/grpcser"
	"pet/iternal/s_reg-ident/jwt/ac"
	"pet/iternal/s_reg-ident/jwt/re"
	"pet/iternal/s_reg-ident/web"
	"pet/pkg/pars"
	"pet/pkg/sql/mysqlcon"
)

var (
	block = make(chan int)

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

	keyref, err := re.GeneratingEncryptionKeys()
	if err != nil {
		log.Fatal(err)
	}
	keyref.Id = &idjwtref

	keyacc, err := ac.GenerateRSAKey()
	if err != nil {
		log.Fatal(err)
	}
	keyacc.Id = &idjwtacc

	hesh := pars.New("reg", "auth", "regstat", "hi")
	hesh.LoadHash(*pathDirUi)

	con := web.Connect{
		MySQL:     dbcon,
		KeyRef:    keyref,
		KeyAcc:    keyacc,
		KeshTempl: &hesh,
	}

	grpcser.StartServerGRPC(*addrGRPC, keyacc.GetPublicKey())
	con.StartServe(addr)
	log.Println("ALL READY")
	<-block
}
