package salt

import (
	rand2 "math/rand"
	"pet/pkg/convert"
	"time"

	"golang.org/x/crypto/argon2"
)

type Salt struct {
	static  []byte
	dynamic []byte
}

func GenerateSalt() *Salt {
	salt := &Salt{}
	salt.static = []byte("hSbPo?Zz")
	var randNum uint8
	for i := 0; i < 8; i++ {
		rand2.Seed(time.Now().UnixNano())
		randNum = uint8(rand2.Intn(255))
		salt.dynamic = append(salt.dynamic, randNum)
	}
	return salt
}

func CreateSaltAuth(salt string) (saltauth *Salt, err error) {
	s := []byte("hSbPo?Zz")
	d, err := convert.StrToByte(salt)
	if err != nil {
		return nil, err
	}
	saltauth = &Salt{static: s, dynamic: d}
	return saltauth, nil
}

func (salt *Salt) GetDynamicSalt() (dunamicsalt []byte) {
	dunamicsalt = salt.dynamic
	return dunamicsalt
}

func (salt *Salt) GeneraterKey(pasbyte []byte) (kay []byte) {
	kay = argon2.Key(pasbyte, salt.static, 12, 32*256, 4, 32)
	kay = argon2.Key(kay, salt.dynamic, 12, 32*256, 4, 32)
	return kay
}
