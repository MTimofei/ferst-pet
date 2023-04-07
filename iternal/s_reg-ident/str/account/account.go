package account

import "pet/iternal/s_reg-ident/str/salt"

type Account struct {
	id       int
	logname  string
	Key      string
	Saltauth *salt.Salt
}

func New(id int, logname, key, saltin string) (a *Account) {
	s := salt.CreateSaltAuth(saltin)

	a = &Account{
		id:       id,
		logname:  logname,
		Key:      key,
		Saltauth: s,
	}
	return a
}
