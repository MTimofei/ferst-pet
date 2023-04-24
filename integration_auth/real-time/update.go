package realtime

import (
	"crypto/rsa"
	"log"
)

type PublicKey struct {
	Key *rsa.PublicKey
}

func (key *PublicKey) update(publiackey *rsa.PublicKey) {
	key.Key = publiackey
}

func UpdatePublicKey(key *PublicKey, keytransfer chan *rsa.PublicKey) {
	for publickey := range keytransfer {
		key.update(publickey)
		log.Println("UpdatePublicKey", key)
	}
}
