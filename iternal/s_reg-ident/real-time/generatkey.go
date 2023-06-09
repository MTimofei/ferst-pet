package realtime

import (
	"crypto/ecdsa"
	"crypto/rsa"
	"log"
	"pet/iternal/s_reg-ident/jwt/ac"
	"pet/iternal/s_reg-ident/jwt/re"
	"sync"
	"time"
)

func RealTimeGenerateEncryptionKeys(transportrefkey chan *ecdsa.PrivateKey, wg *sync.WaitGroup) {

	defer wg.Done()
	var timesleep time.Duration
	for {
		keyref, err := re.GeneratingEncryptionKeys()
		if err != nil {
			log.Panicln("EncryptionKeys", err)
			timesleep = 5 * time.Second
		} else {
			transportrefkey <- keyref
			timesleep = 1 * time.Hour
		}
		log.Println("RealTimeGenerateEncryptionKeys")
		time.Sleep(timesleep)
	}
}
func RealTimeGenerateRSAKey(transportacckey chan *rsa.PrivateKey, wg *sync.WaitGroup) {

	defer wg.Done()
	var timesleep time.Duration
	for {
		keyref, err := ac.GenerateRSAKey()
		if err != nil {
			log.Panicln("RSAKey", err)
			timesleep = 5 * time.Second
		} else {
			transportacckey <- keyref
			timesleep = 5 * time.Minute
		}
		log.Println("RealTimeGenerateRSAKey")
		time.Sleep(timesleep)
	}

}
