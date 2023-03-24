package salt

import (
	"fmt"
	"golang.org/x/crypto/argon2"
	"log"
	rand2 "math/rand"
	"time"
)

type Salt struct {
	static  []byte
	dynamic []byte
}

func GenerateSalt() *Salt {
	log.Println("21data transmission")
	salt := &Salt{}
	salt.static = []byte("hSbPo?Zz")
	var randNum uint8
	log.Println("22data transmission")
	for i := 0; i < 8; i++ {
		rand2.Seed(time.Now().UnixNano())
		randNum = uint8(rand2.Intn(255))
		fmt.Println(randNum)
		salt.dynamic = append(salt.dynamic, randNum)
	}
	log.Println("23data transmission")
	return salt
}
func (salt *Salt) GetDyanmicSalt() (dunamicsalt []byte) {
	dunamicsalt = salt.dynamic
	return dunamicsalt
}

func (salt *Salt) GeneraterKey(pasbyte []byte) (kay []byte) {
	kay = argon2.Key(pasbyte, salt.static, 12, 32*256, 4, 32)
	kay = argon2.Key(kay, salt.dynamic, 12, 32*256, 4, 32)
	return kay
}
