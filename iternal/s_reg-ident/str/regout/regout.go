package regout

import (
	"pet/iternal/s_reg-ident/str/regin"
	"pet/iternal/s_reg-ident/str/salt"
)

type RegDataDB struct {
	loginname       string
	email           string
	passwordhashStr string
	passwordhashBt  []byte
	saltdynamicStr  string
	saltdynamicBt   []byte
}

func NewInToOut(salt *salt.Salt, rdin *regin.RegDataIn, key []byte) (rdout *RegDataDB) {
	var intermediateint []int
	var intermediatestr []string

	rdout.loginname, rdout.email = rdin.GetRDIn()
	rdout.passwordhashBt = key
	rdout.saltdynamicBt = salt.GetDyanmicSalt()

	for i := 0; i < len(rdout.passwordhashBt); i++ {

	}

	return rdout
}
