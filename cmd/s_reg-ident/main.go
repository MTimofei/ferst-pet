package main

//accaunt_ser fZLma7
import (
	"flag"
	"log"
	"net/http"
	"pet/iternal/s_reg-ident/jwt/re"
	"pet/iternal/s_reg-ident/web"
	"pet/pkg/pars"
	"pet/pkg/sql/mysqlcon"
)

var (
	idjwt     = 0
	addr      = flag.String("addr", "localhost:8889", "adderss server")
	addrMySQL = flag.String("adder-MySQL", "accaunt_ser:fZLma7@/ppa?parseTime=true", "adderss mysql")
)

func main() {
	flag.Parse()
	dbcon, err := mysqlcon.OpenMySQLDB(addrMySQL)
	if err != nil {
		log.Fatal(err)
	}
	defer dbcon.Close()
	k, err := re.GeneratingEncryptionKeys()
	if err != nil {
		log.Fatal(err)
	}
	k.Id = &idjwt
	hesh := pars.New("reg", "auth", "regstat")
	hesh.LoadHash("ui/HTML/")
	con := web.Connect{
		MySQL:     dbcon,
		K:         k,
		HashTempl: &hesh,
	}

	err = http.ListenAndServe(*addr, con.Router())
	if err != nil {
		log.Fatal(err)
	}
}
