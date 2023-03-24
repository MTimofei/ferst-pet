package salt

import rand2 "math/rand"

type Salt struct {
	static  []byte
	dynamic []byte
}

func GenerateSalt() Salt {
	var salt Salt
	salt.static = []byte("")
	salt.dynamic = []byte(string(rand2.Uint32()))
	return salt
}
