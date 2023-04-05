package comsql

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"pet/iternal/s_reg-ident/str/regin"
	"pet/iternal/s_reg-ident/str/salt"
	"pet/pkg/convert"
)

func CheckUinquenessLogin(db *sql.DB, r *http.Request) (err error) {
	query := fmt.Sprintf(`SELECT log_name FROM account WHERE log_name = '%s';`, r.FormValue("name"))
	ra, err := db.Query(query)
	if err != nil {
		return err
	}
	//fmt.Println(ra.Next())
	if ra.Next() {
		return errors.New("login already exists")
	} else {
		return nil
	}
}

func CheckUinquenessEmail(db *sql.DB, r *http.Request) (err error) {
	query := fmt.Sprintf(`SELECT email FROM account WHERE email = '%s';`, r.FormValue("email"))
	ra, err := db.Query(query)
	if err != nil {
		return err
	}
	if ra.Next() {
		return errors.New("email already exists")
	} else {
		return nil
	}
}

func SendRegData(db *sql.DB, salt *salt.Salt, rdin *regin.RegDataIn, key []byte) error {
	l, e := rdin.GetRDIn()
	k := convert.IntToStr(convert.ByteToInt(key))
	s := convert.IntToStr(convert.ByteToInt(salt.GetDyanmicSalt()))
	query := fmt.Sprintf(`INSERT INTO account (log_name, email, password_hash, salt_hash) VALUES ('%s', '%s', '%s', '%s');`,
		l, e, k, s)
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
