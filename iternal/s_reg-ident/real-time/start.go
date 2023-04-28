package realtime

import (
	"crypto/ecdsa"
	"crypto/rsa"
	"pet/iternal/s_reg-ident/web"
)

func StartUpdataKey(block chan int, con *web.Connect, transportrefkey chan *ecdsa.PrivateKey, transportacckey chan *rsa.PrivateKey) {
	go RealTimeGenerateEncryptionKeys(transportrefkey)
	go RealTimeGenerateRSAKey(transportacckey)

	go UpdateRefPrivateKey(con, transportrefkey)
	go UpdateAccPrivateKey(con, transportacckey)
	<-block
}
