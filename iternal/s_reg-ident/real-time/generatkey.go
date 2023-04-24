package realtime

import (
	"crypto/ecdsa"
	"crypto/rsa"
	"log"
	"pet/iternal/s_reg-ident/jwt/ac"
	"pet/iternal/s_reg-ident/jwt/re"
	"time"
)

func RealTimeGenerateEncryptionKeys(transportrefkey chan *ecdsa.PrivateKey) {
	go func() {
		var timesleep time.Duration
		for {
			keyref, err := re.GeneratingEncryptionKeys()
			if err != nil {
				log.Panicln("EncryptionKeys", err)
			} else {
				transportrefkey <- keyref
				timesleep = 10 * time.Second
			}
			log.Println("RealTimeGenerateEncryptionKeys", *keyref)
			time.Sleep(timesleep)
		}
	}()
}
func RealTimeGenerateRSAKey(transportacckey chan *rsa.PrivateKey) {
	go func() {
		var timesleep time.Duration
		for {
			keyref, err := ac.GenerateRSAKey()
			if err != nil {
				log.Panicln("RSAKey", err)
			} else {
				transportacckey <- keyref
				timesleep = 10 * time.Second
			}
			log.Println("RealTimeGenerateRSAKey", *keyref)
			time.Sleep(timesleep)
		}
	}()
}
