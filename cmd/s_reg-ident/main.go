package main

//accaunt_ser fZLma7
import (
	"flag"
	"log"
	"net/http"
	"pet/iternal/s_reg-ident/jwt/ac"
	"pet/iternal/s_reg-ident/jwt/re"
	"pet/iternal/s_reg-ident/web"
	"pet/pkg/pars"
	"pet/pkg/sql/mysqlcon"
)

var (
	idjwtref int64 = 0
	idjwtacc int64 = 0

	addr      = flag.String("addr", "localhost:8889", "adderss server")
	pathDirUi = flag.String("puth-dir-ui", "ui/HTML/", "puth ui directory")
	addrMySQL = flag.String("adder-MySQL", "accaunt_ser:fZLma7@/ppa?parseTime=true", "adderss mysql")
)

func main() {
	flag.Parse()

	dbcon, err := mysqlcon.OpenMySQLDB(addrMySQL)
	if err != nil {
		log.Fatal(err)
	}
	defer dbcon.Close()

	kref, err := re.GeneratingEncryptionKeys()
	if err != nil {
		log.Fatal(err)
	}
	kref.Id = &idjwtref

	kacc, err := ac.GenerateRSAKey()
	if err != nil {
		log.Fatal(err)
	}
	kacc.Id = &idjwtacc

	hesh := pars.New("reg", "auth", "regstat", "hi")
	hesh.LoadHash(*pathDirUi)

	con := web.Connect{
		MySQL:     dbcon,
		KRef:      kref,
		KAcc:      kacc,
		HashTempl: &hesh,
	}

	err = http.ListenAndServe(*addr, con.Router())
	if err != nil {
		log.Fatal(err)
	}
}
