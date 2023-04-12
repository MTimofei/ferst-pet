package main

//accaunt_ser fZLma7
import (
	"flag"
	"log"
	"net/http"
	"pet/iternal/s_reg-ident/jwt/re"
	"pet/iternal/s_reg-ident/web"
	"pet/pkg/sql/mysqlcon"
)

var (
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
	con := web.Connect{
		MySQL: dbcon,
		K:     k,
	}

	err = http.ListenAndServe(*addr, con.Router())
	if err != nil {
		log.Fatal(err)
	}
}
