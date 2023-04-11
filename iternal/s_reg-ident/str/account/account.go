package account

import "pet/iternal/s_reg-ident/str/salt"

type Account struct {
	Id       int
	Logname  string
	Key      string
	Saltauth *salt.Salt
}

func New(id int, logname, key, saltin string) (a *Account) {
	s := salt.CreateSaltAuth(saltin)

	a = &Account{
		Id:       id,
		Logname:  logname,
		Key:      key,
		Saltauth: s,
	}
	return a
}
