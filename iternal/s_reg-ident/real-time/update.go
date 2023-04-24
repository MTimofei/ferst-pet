package realtime

import (
	"crypto/ecdsa"
	"crypto/rsa"
	"log"
	"pet/iternal/s_reg-ident/web"
)

func UpdateRefPrivateKey(con *web.Connect, transportrefkey chan *ecdsa.PrivateKey) {
	for privatekey := range transportrefkey {
		con.KeyRef.Update(privatekey)
		log.Println("updateRefPrivateKey")
	}
}

func UpdateAccPrivateKey(con *web.Connect, transportacckey chan *rsa.PrivateKey) {
	for privatekey := range transportacckey {
		con.KeyAcc.Update(privatekey)
		log.Println("updateAccPrivateKey")
	}
}
