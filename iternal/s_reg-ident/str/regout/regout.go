package regout

import (
	"pet/pkg/convert"

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

func NewReg(salt *salt.Salt, rdin *regin.RegDataIn, key []byte) (rdout *RegDataDB) {

	rdout.loginname, rdout.email = rdin.GetRDIn()
	rdout.passwordhashBt = key
	rdout.saltdynamicBt = salt.GetDyanmicSalt()
	rdout.passwordhashStr = convert.IntToStr(convert.ByteToInt(rdout.passwordhashBt))
	rdout.saltdynamicStr = convert.IntToStr(convert.ByteToInt(rdout.saltdynamicBt))

	return rdout
}
