package realtime

import (
	"crypto/ecdsa"
	"crypto/rsa"
	"log"
	"pet/iternal/s_reg-ident/web"
	"sync"
)

func UpdateRefPrivateKey(con *web.Connect, transportrefkey chan *ecdsa.PrivateKey, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	for privatekey := range transportrefkey {
		con.KeyRef.Update(privatekey)
		log.Println("updateAccPrivateKey")
	}
}

func UpdateAccPrivateKey(con *web.Connect, transportacckey chan *rsa.PrivateKey, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	for privatekey := range transportacckey {
		con.KeyAcc.Update(privatekey)
		log.Println("updateAccPrivateKey")
	}
}
