package realtime

import (
	"crypto/ecdsa"
	"crypto/rsa"
	"pet/iternal/s_reg-ident/web"
	"sync"
)

func StartUpdataKey(con *web.Connect, transportrefkey chan *ecdsa.PrivateKey, transportacckey chan *rsa.PrivateKey) {
	wg := &sync.WaitGroup{}
	wg.Add(4)
	go RealTimeGenerateEncryptionKeys(transportrefkey, wg)
	go RealTimeGenerateRSAKey(transportacckey, wg)

	go UpdateRefPrivateKey(con, transportrefkey, wg)
	go UpdateAccPrivateKey(con, transportacckey, wg)

	wg.Wait()
}
