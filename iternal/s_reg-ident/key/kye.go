package key

import (
	"golang.org/x/crypto/argon2"
	rand2 "math/rand"
)

func GenerateDenamicSolt() []byte {
	return []byte(string(rand2.Uint32()))
}

func GeneraterKey(dynamicsalt []byte, pasbyte []byte) (kay []byte) {
	statsalt := []byte("0EFnX)34")
	kay = argon2.Key(pasbyte, statsalt, 12, 32*256, 4, 32)
	kay = argon2.Key(kay, dynamicsalt, 12, 32*256, 4, 32)
	return kay
}
