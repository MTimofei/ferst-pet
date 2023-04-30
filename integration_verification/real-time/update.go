package realtime

import (
	"crypto/rsa"
	"log"
	"pet/integration_verification/real-time/grpcclient"
	"sync"
)

type PublicKey struct {
	Key *rsa.PublicKey
}

func StartUpdate(addrGRPC *string, key *PublicKey, keytransfer chan *rsa.PublicKey) {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go grpcclient.RealTimeGetKyeМiaGRPC(addrGRPC, keytransfer, wg)
	go updatePublicKey(key, keytransfer, wg)
	wg.Wait()
}

func (key *PublicKey) update(publiackey *rsa.PublicKey) {
	key.Key = publiackey
}

func updatePublicKey(key *PublicKey, keytransfer chan *rsa.PublicKey, wg *sync.WaitGroup) {
	defer wg.Done()
	for publickey := range keytransfer {
		key.update(publickey)
		log.Println("UpdatePublicKey")
	}
}
