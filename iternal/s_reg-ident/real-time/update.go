package realtime

import (
	"crypto/ecdsa"
	"crypto/rsa"
	"pet/iternal/s_reg-ident/web"
)

func RealTimeUpdatePrivateKey(con *web.Connect, transportrefkey chan *ecdsa.PrivateKey, transportacckey chan *rsa.PrivateKey) {
	go UpdateRefPrivateKey(con, transportrefkey)
	go UpdateAccPrivateKey(con, transportacckey)
}

func UpdateRefPrivateKey(con *web.Connect, transportrefkey chan *ecdsa.PrivateKey) {
	for privatekey := range transportrefkey {
		con.KeyRef.Update(privatekey)
	}
}

func UpdateAccPrivateKey(con *web.Connect, transportacckey chan *rsa.PrivateKey) {
	for privatekey := range transportacckey {
		con.KeyAcc.Update(privatekey)
	}
}
