package realtime

import (
	"crypto/rsa"
	"log"
	"pet/integration_verification/real-time/grpcclient"
)

type PublicKey struct {
	Key *rsa.PublicKey
}

func StartUpdate(addrGRPC *string, key *PublicKey, keytransfer chan *rsa.PublicKey, blok chan int) {
	go grpcclient.RealTimeGetKyeМiaGRPC(addrGRPC, keytransfer)
	go updatePublicKey(key, keytransfer)
	<-blok
}

func (key *PublicKey) update(publiackey *rsa.PublicKey) {
	key.Key = publiackey
}

func updatePublicKey(key *PublicKey, keytransfer chan *rsa.PublicKey) {
	for publickey := range keytransfer {
		key.update(publickey)
		log.Println("UpdatePublicKey", key)
	}
}
