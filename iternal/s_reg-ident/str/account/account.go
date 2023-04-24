package account

import "pet/iternal/s_reg-ident/str/salt"

type Account struct {
	Id           int
	Logname      string
	HashPassword string
	Saltauth     *salt.Salt
}

func New(id int, logname, heshpassword, saltin string) (a *Account) {
	s := salt.CreateSaltAuth(saltin)

	a = &Account{
		Id:           id,
		Logname:      logname,
		HashPassword: heshpassword,
		Saltauth:     s,
	}
	return a
}
