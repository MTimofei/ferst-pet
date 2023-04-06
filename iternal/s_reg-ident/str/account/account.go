package account

import "pet/iternal/s_reg-ident/str/salt"

type Account struct {
	id       int
	logname  string
	Key      string
	Saltauth *salt.Salt
}

func New(id int, logname, key, saltin string) (a *Account, err error) {
	s, err := salt.CreateSaltAuth(saltin)
	if err != nil {
		return nil, err
	}
	a = &Account{
		id:       id,
		logname:  logname,
		Key:      key,
		Saltauth: s,
	}
	return a, nil
}
