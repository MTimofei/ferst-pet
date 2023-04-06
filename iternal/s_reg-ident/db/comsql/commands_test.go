package comsql_test

import (
	"fmt"
	"net/http"
	"net/url"
	"pet/iternal/s_reg-ident/db/comsql"
	"pet/iternal/s_reg-ident/str/account"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestCheckUinquessLogin(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating sqlmock DB: %v", err)
	}
	defer db.Close()

	name := &[]struct {
		numtest     string
		dataDB      string
		dataclient  string
		correctresp error
	}{
		{numtest: "test_1", dataDB: "mark", dataclient: "mark", correctresp: fmt.Errorf("login already exists")},
		{numtest: "test_2", dataDB: "", dataclient: "bob", correctresp: nil},
		{numtest: "test_3", dataDB: "bob", dataclient: "bob", correctresp: fmt.Errorf("login already exists")},
		{numtest: "test_4", dataDB: "les", dataclient: "les", correctresp: fmt.Errorf("login already exists")},
		{numtest: "test_5", dataDB: "", dataclient: "tim", correctresp: nil},
	}
	for _, n := range *name {
		t.Run(n.numtest, func(t *testing.T) {
			rows := sqlmock.NewRows([]string{"logname"})
			if n.dataDB != "" {
				rows.AddRow(n.dataDB)
			}

			query := fmt.Sprintf(`SELECT log_name FROM account WHERE log_name = '%s';`, n.dataclient)
			mock.ExpectQuery(query).WillReturnRows(rows)

			req, err := http.NewRequest("POST", "/test", nil)
			if err != nil {
				t.Fatalf("failed to create request: %v", err)
			}

			req.Form = url.Values{}
			req.Form.Add("name", n.dataclient)

			resolt := comsql.CheckUinquenessLogin(db, req)

			if err = mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}

			if !reflect.DeepEqual(resolt, n.correctresp) {
				t.Errorf("\nbd log_name:	%s\nclient log_name:	%s\nresolt:	 %s", n.dataDB, n.dataclient, resolt)
			}
		})
	}
}

func TestCheckUinquessEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating sqlmock DB: %v", err)
	}
	defer db.Close()

	name := &[]struct {
		numtest     string
		dataDB      string
		dataclient  string
		correctresp error
	}{
		{numtest: "test_1", dataDB: "mark@test.com", dataclient: "mark@test.com", correctresp: fmt.Errorf("email already exists")},
		{numtest: "test_2", dataDB: "", dataclient: "bob@test.com", correctresp: nil},
		{numtest: "test_3", dataDB: "bob@test.com", dataclient: "bob@test.com", correctresp: fmt.Errorf("email already exists")},
		{numtest: "test_4", dataDB: "les@test.com", dataclient: "les@test.com", correctresp: fmt.Errorf("email already exists")},
		{numtest: "test_5", dataDB: "", dataclient: "tim@test.com", correctresp: nil},
	}
	for _, n := range *name {
		t.Run(n.numtest, func(t *testing.T) {
			rows := sqlmock.NewRows([]string{"email"})
			if n.dataDB != "" {
				rows.AddRow(n.dataDB)
			}

			query := fmt.Sprintf(`SELECT email FROM account WHERE email = '%s';`, n.dataclient)
			mock.ExpectQuery(query).WillReturnRows(rows)

			req, err := http.NewRequest("POST", "/test", nil)
			if err != nil {
				t.Fatalf("failed to create request: %v", err)
			}

			req.Form = url.Values{}
			req.Form.Add("email", n.dataclient)

			resolt := comsql.CheckUinquenessEmail(db, req)

			if err = mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}

			if !reflect.DeepEqual(resolt, n.correctresp) {
				t.Errorf("\nbd log_name:	%s\nclient log_name:	%s\nresolt:	 %s", n.dataDB, n.dataclient, resolt)
			}
		})
	}
}

func TestGetAccountData(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating sqlmock DB: %v", err)
	}
	defer db.Close()

	type Resp struct {
		acc *account.Account
		err error
	}

	type DataDB struct {
		id       int
		password string
		salt     string
	}

	account := &[]struct {
		numtest            string
		db                 DataDB
		dataclient_logname string
		correctresp        Resp
	}{
		{numtest: "test_1", db: DataDB{id: 1, password: "001", salt: "00"}, dataclient_logname: "i", correctresp: Resp{acc: account.New(1, "i", "001", "00"), err: nil}},
		{numtest: "test_2", db: DataDB{id: 2, password: "112", salt: "00"}, dataclient_logname: "o", correctresp: Resp{acc: account.New(2, "o", "112", "00"), err: nil}},
		{numtest: "test_3", db: DataDB{id: 3, password: "234", salt: "00"}, dataclient_logname: "p", correctresp: Resp{acc: account.New(3, "p", "234", "00"), err: nil}},
		{numtest: "test_4", db: DataDB{id: 4, password: "24", salt: "00"}, dataclient_logname: "u", correctresp: Resp{acc: account.New(4, "u", "24", "00"), err: nil}},
	}

	for _, n := range *account {

		t.Run(n.numtest, func(t *testing.T) {

			rows := sqlmock.NewRows([]string{"id", "password", "salt"})
			rows.AddRow(n.db.id, n.db.password, n.db.salt)

			query := fmt.Sprintf(`SELECT id, password_hash, salt_hash FROM account WHERE log_name = '%s';`, n.dataclient_logname)
			mock.ExpectQuery(query).WillReturnRows(rows)

			resolt := &Resp{}
			resolt.acc, resolt.err = comsql.GetAccountData(db, n.dataclient_logname)

			if err = mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
			if resolt.err != n.correctresp.err {
				t.Errorf("Unexpected error: %v", resolt.err)
			}
			if reflect.DeepEqual(resolt.acc, *n.correctresp.acc) {
				t.Errorf("Unexpected account data. Expected: %v, got: %v", n.correctresp.acc, resolt.acc)
			}
		})
	}
}
